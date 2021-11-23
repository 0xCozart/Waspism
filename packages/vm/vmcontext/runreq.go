package vmcontext

import (
	"errors"
	"math"
	"math/big"
	"runtime/debug"
	"time"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/vmcontext/vmtxbuilder"
	"golang.org/x/xerrors"
)

// RunTheRequest processes each iscp.RequestData in the batch
func (vmctx *VMContext) RunTheRequest(req iscp.RequestData, requestIndex uint16) error {
	// prepare context for the request
	vmctx.req = req
	vmctx.requestIndex = requestIndex
	vmctx.requestEventIndex = 0
	vmctx.entropy = hashing.HashData(vmctx.entropy[:])
	vmctx.callStack = vmctx.callStack[:0]

	if err := vmctx.earlyCheckReasonToSkip(); err != nil {
		return err
	}
	vmctx.loadChainConfig()
	vmctx.locateTargetContract()

	// at this point state update is empty
	// so far there were no panics except optimistic reader
	// No prepare state update (buffer) for mutations and panics

	txsnapshot := vmctx.createTxBuilderSnapshot()
	vmctx.currentStateUpdate = state.NewStateUpdate(vmctx.virtualState.Timestamp().Add(1 * time.Nanosecond))

	// catches error which is not the request or contract fault
	// If it occurs, the request is just skipped
	err := util.CatchPanicReturnError(func() {
		// transfer all attached assets to the sender's account
		vmctx.creditAssetsToChain()
		// load gas and fee policy, calculate and set gas budget
		vmctx.prepareGasBudget()
		// run the contract program
		vmctx.callTheContract()
	}, vmtxbuilder.ErrInputLimitExceeded, vmtxbuilder.ErrOutputLimitExceeded)

	if err != nil {
		// transaction limits exceeded. Skipping the request. Rollback
		vmctx.restoreTxBuilderSnapshot(txsnapshot)
		vmctx.currentStateUpdate = nil
		return err
	}
	vmctx.virtualState.ApplyStateUpdates(vmctx.currentStateUpdate)
	vmctx.currentStateUpdate = nil
	return nil
}

// creditAssetsToChain credits L1 accounts with attached assets and accrues all of them to the sender's account on-chain
func (vmctx *VMContext) creditAssetsToChain() {
	if vmctx.req.IsOffLedger() {
		// off ledger requests does not bring any deposit
		return
	}
	// ---- update transaction builder
	vmctx.txbuilder.ConsumeOutput(vmctx.req)
	vmctx.txbuilder.AddDeltaIotas(vmctx.req.Assets().Iotas)
	for _, nt := range vmctx.req.Assets().Tokens {
		vmctx.txbuilder.AddDeltaNativeToken(nt.ID, nt.Amount)
	}
	// ---- end update transaction builder

	// ---- update the state, the account ledger
	// NOTE: sender account will be CommonAccount if sender address is not available
	vmctx.creditToAccount(vmctx.req.SenderAccount(), vmctx.req.Assets())
	// ---- end update state

	// here transaction builder must be consistent itself and be consistent with the state (the accounts)

	_, _, isBalanced := vmctx.txbuilder.TotalAssets()
	if !isBalanced {
		panic("internal inconsistency: transaction builder is not balanced")
	}
	// TODO check if total assets are consistent with the state
}

func (vmctx *VMContext) prepareGasBudget() {
	vmctx.loadGasPolicy()
	vmctx.calculateAffordableGasBudget()
	vmctx.gasSetBudget(vmctx.gasBudgetAffordable)
}

// callTheContract runs the contract. It catches and processes all panics except the one which cancel the whole block
func (vmctx *VMContext) callTheContract() {
	// TODO
	txsnapshot := vmctx.createTxBuilderSnapshot()
	snapMutations := vmctx.currentStateUpdate.Clone()

	if vmctx.req.IsOffLedger() {
		vmctx.updateOffLedgerRequestMaxAssumedNonce()
	}
	vmctx.lastError = nil
	func() {
		defer func() {
			vmctx.lastError = checkVMPluginPanic()
			if vmctx.lastError == nil {
				return
			}
			vmctx.lastResult = nil
			vmctx.Debugf("%v", vmctx.lastError)
			vmctx.Debugf(string(debug.Stack()))
		}()
		vmctx.callFromRequest()
	}()
	if vmctx.lastError != nil {
		// panic happened during VM plugin call
		// restore the state
		vmctx.restoreTxBuilderSnapshot(txsnapshot)
		vmctx.currentStateUpdate = snapMutations
	}
	vmctx.chargeGasFee()
}

func checkVMPluginPanic() error {
	r := recover()
	if r == nil {
		return nil
	}
	// re-panic-ing if error it not user nor VM plugin fault.
	// Otherwise, the panic is wrapped into the returned error, including gas-related panic
	switch err := r.(type) {
	case *kv.DBError:
		panic(err)
	case error:
		if errors.Is(err, coreutil.ErrorStateInvalidated) {
			panic(err)
		}
		if errors.Is(err, vmtxbuilder.ErrOutputLimitExceeded) {
			panic(err)
		}
		if errors.Is(err, vmtxbuilder.ErrInputLimitExceeded) {
			panic(err)
		}
	}
	return xerrors.Errorf("exception: %w", r)
}

// callFromRequest is the call itself. Assumes sc exists
func (vmctx *VMContext) callFromRequest() {
	vmctx.Debugf("callFromRequest: %s", vmctx.req.ID().String())

	// calling only non view entry points. Calling the view will trigger error and fallback
	entryPoint := vmctx.req.Target().EntryPoint
	targetContract := vmctx.contractRecord.Hname()
	vmctx.lastResult, vmctx.lastError = vmctx.callNonViewByProgramHash(
		targetContract,
		entryPoint,
		vmctx.req.Params(),
		vmctx.req.Transfer(),
		vmctx.contractRecord.ProgramHash,
	)
}

// chargeGasFee takes burned tokens from the sender's account
// It should always be enough because gas budget is set affordable
func (vmctx *VMContext) chargeGasFee() {
	if vmctx.req.SenderAddress() == nil {
		panic("inconsistency: vmctx.req.Request().SenderAddress() == nil")
	}
	tokensToMove := vmctx.GasBurned() / vmctx.gasPolicyGasPerGasToken
	transferToValidator := &iscp.Assets{}
	if vmctx.gasFeeTokenNotIota {
		transferToValidator.Tokens = iotago.NativeTokens{{vmctx.gasFeeTokenID, new(big.Int).SetUint64(tokensToMove)}}
	} else {
		transferToValidator.Iotas = tokensToMove
	}
	sender := vmctx.req.SenderAccount()
	// TODO split validator/chain owner
	vmctx.mustMoveBetweenAccounts(sender, vmctx.task.ValidatorFeeTarget, transferToValidator)
}

// calculateAffordableGasBudget checks the account of the sender and calculates affordable gas budget
// Affordable gas budget is calculated from gas budget provided in the request by the user and taking into account
// how many tokens the sender has in its account.
// Safe arithmetics is used
func (vmctx *VMContext) calculateAffordableGasBudget() {
	if vmctx.req.SenderAddress() == nil {
		panic("inconsistency: vmctx.req.SenderAddress() == nil")
	}
	tokensAvailable := uint64(0)
	if vmctx.gasFeeTokenNotIota {
		tokensAvailableBig := vmctx.GetTokenBalance(vmctx.req.SenderAccount(), &vmctx.gasFeeTokenID)
		if tokensAvailableBig != nil {
			// safely subtract the transfer from the sender to the target
			if transfer := vmctx.req.Transfer(); transfer != nil {
				if transferTokens := iscp.FindNativeTokenBalance(transfer.Tokens, &vmctx.gasFeeTokenID); transferTokens != nil {
					if tokensAvailableBig.Cmp(transferTokens) < 0 {
						tokensAvailableBig.SetUint64(0)
					} else {
						tokensAvailableBig.Sub(tokensAvailableBig, transferTokens)
					}
				}
			}
			if tokensAvailableBig.IsUint64() {
				tokensAvailable = tokensAvailableBig.Uint64()
			} else {
				tokensAvailable = math.MaxUint64
			}
		}
	} else {
		tokensAvailable = vmctx.GetIotaBalance(vmctx.req.SenderAccount())
		// safely subtract the transfer from the sender to the target
		if transfer := vmctx.req.Transfer(); transfer != nil {
			if tokensAvailable < transfer.Iotas {
				tokensAvailable = 0
			} else {
				tokensAvailable -= transfer.Iotas
			}
		}
	}
	if tokensAvailable < math.MaxUint64/vmctx.gasPolicyGasPerGasToken {
		vmctx.gasBudgetAffordable = tokensAvailable * vmctx.gasPolicyGasPerGasToken
	} else {
		vmctx.gasBudgetAffordable = math.MaxUint64
	}

	// TODO introduce minimum balance on account ?
	vmctx.gasBudgetFromRequest = vmctx.req.GasBudget()
	vmctx.gasBudget = vmctx.gasBudgetFromRequest
	if vmctx.gasBudget > vmctx.gasBudgetAffordable {
		vmctx.gasBudget = vmctx.gasBudgetAffordable
	}
}

func (vmctx *VMContext) loadGasPolicy() {
	// TODO load from governance contract
	vmctx.gasFeeTokenNotIota = false
	vmctx.gasFeeTokenID = iotago.NativeTokenID{}
	vmctx.gasPolicyFixedBudget = false
	vmctx.gasPolicyGasPerGasToken = 100
}

func (vmctx *VMContext) locateTargetContract() {
	// find target contract
	targetContract := vmctx.req.Target().Contract
	var ok bool
	vmctx.contractRecord, ok = vmctx.findContractByHname(targetContract)
	if !ok {
		vmctx.Warnf("contract not found: %s", targetContract)
	}
	if vmctx.contractRecord.Hname() == 0 {
		vmctx.Warnf("default contract will be called")
	}
}

// loadChainConfig only makes sense if chain is already deployed
func (vmctx *VMContext) loadChainConfig() {
	if vmctx.isInitChainRequest() {
		vmctx.chainOwnerID = vmctx.req.SenderAccount()
		return
	}
	cfg := vmctx.getChainInfo()
	vmctx.chainOwnerID = cfg.ChainOwnerID
	vmctx.maxEventSize = cfg.MaxEventSize
	vmctx.maxEventsPerReq = cfg.MaxEventsPerReq
	//vmctx.feeColor, vmctx.ownerFee, vmctx.validatorFee = vmctx.getFeeInfo()  // TODO fee policy
}

func (vmctx *VMContext) isInitChainRequest() bool {
	target := vmctx.req.Target()
	return target.Contract == root.Contract.Hname() && target.EntryPoint == iscp.EntryPointInit
}
