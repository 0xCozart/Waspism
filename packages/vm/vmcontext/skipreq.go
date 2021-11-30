package vmcontext

import (
	"time"

	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/blocklog"
	"github.com/iotaledger/wasp/packages/vm/vmcontext/vmtxbuilder"
	"golang.org/x/xerrors"
)

const (
	// OffLedgerNonceStrictOrderTolerance how many steps back the nonce is considered too old
	// within this limit order of nonces is not checked
	OffLedgerNonceStrictOrderTolerance = 10000

	// OnLedgerTooOldMilestonesBack how many milestones back confirmed UTXO we consider too old
	OnLedgerTooOldMilestonesBack = 10

	// ExpiryUnlockSafetyWindowDuration creates safety window around time assumption,
	// the UTXO won't be consumed to avoid race conditions
	ExpiryUnlockSafetyWindowDuration  = 1 * time.Minute
	ExpiryUnlockSafetyWindowMilestone = 3
)

// earlyCheckReasonToSkip checks if request must be ignored without even modifying the state
func (vmctx *VMContext) earlyCheckReasonToSkip() error {
	var err error
	if vmctx.req.IsOffLedger() {
		err = vmctx.checkReasonToSkipOffLedger()
	} else {
		err = vmctx.checkReasonToSkipOnLedger()
	}
	return err
}

// checkReasonRequestProcessed checks if request ID is already in the blocklog
func (vmctx *VMContext) checkReasonRequestProcessed() error {
	vmctx.pushCallContext(blocklog.Contract.Hname(), nil, nil)
	defer vmctx.popCallContext()

	reqid := vmctx.req.ID()
	if blocklog.MustIsRequestProcessed(vmctx.State(), &reqid) {
		return xerrors.New("already processed")
	}
	return nil
}

// checkReasonToSkipOffLedger checks reasons to skip off ledger request
func (vmctx *VMContext) checkReasonToSkipOffLedger() error {
	// first checks if it is already in backlog
	if err := vmctx.checkReasonRequestProcessed(); err != nil {
		return err
	}
	vmctx.pushCallContext(accounts.Contract.Hname(), nil, nil)
	defer vmctx.popCallContext()

	// check the account. It must exist
	// TODO optimize: check the account balances and fetch nonce in one call
	// off-ledger account must exist, i.e. it should have non zero balance on the chain
	if _, exists := accounts.GetAccountAssets(vmctx.State(), vmctx.req.SenderAccount()); !exists {
		// TODO check minimum balance. Require some minimum balance
		return xerrors.Errorf("unverified account for off-ledger request: %s", vmctx.req.SenderAccount())
	}

	// this is a replay protection measure for off-ledger requests assuming in the batch order of requests is random.
	// It is checking if nonce is not too old. See replay-off-ledger.md
	maxAssumed := accounts.GetMaxAssumedNonce(vmctx.State(), vmctx.req.SenderAddress())

	nonce := vmctx.req.Unwrap().OffLedger().Nonce()
	vmctx.Debugf("vmctx.validateRequest - nonce check - maxAssumed: %d, tolerance: %d, request nonce: %d ",
		maxAssumed, OffLedgerNonceStrictOrderTolerance, nonce)

	if maxAssumed < OffLedgerNonceStrictOrderTolerance {
		return nil
	}
	if nonce > maxAssumed-OffLedgerNonceStrictOrderTolerance {
		return xerrors.Errorf("nonce %d is too old", nonce)
	}
	return nil
}

// checkReasonToSkipOnLedger check reasons to skip UTXO request
func (vmctx *VMContext) checkReasonToSkipOnLedger() error {
	if err := vmctx.checkReasonReturnAmount(); err != nil {
		return err
	}
	if err := vmctx.checkReasonTimeLock(); err != nil {
		return err
	}
	if err := vmctx.checkReasonExpiry(); err != nil {
		return err
	}
	if vmctx.txbuilder.InputsAreFull() {
		return vmtxbuilder.ErrInputLimitExceeded
	}
	if err := vmctx.checkReasonRequestProcessed(); err != nil {
		return err
	}
	// if the output was not consumed during last OnLedgerTooOldMilestonesBack milestones, skip it
	if vmctx.req.Unwrap().UTXO().Metadata().MilestoneIndex < vmctx.task.TimeAssumption.MilestoneIndex-OnLedgerTooOldMilestonesBack {
		return xerrors.Errorf("more than %d milestones back", OnLedgerTooOldMilestonesBack)
	}
	return nil
}

// checkReasonTimeLock checking timelock conditions based on time assumptions.
// VM must ensure that the UTXO can be unlocked
func (vmctx *VMContext) checkReasonTimeLock() error {
	lock := vmctx.req.Unwrap().UTXO().Features().TimeLock()
	if lock != nil {
		if lock.Time.Before(vmctx.finalStateTimestamp) {
			return xerrors.Errorf("can't be consumed due to lock until %v", vmctx.finalStateTimestamp)
		}
		if lock.MilestoneIndex != 0 && vmctx.task.TimeAssumption.MilestoneIndex < lock.MilestoneIndex {
			return xerrors.Errorf("can't be consumed due to lock until milestone index #%v", vmctx.task.TimeAssumption.MilestoneIndex)
		}
	}
	return nil
}

// checkReasonExpiry checking expiry conditions based on time assumptions.
// VM must ensure that the UTXO can be unlocked
func (vmctx *VMContext) checkReasonExpiry() error {
	expiry, senderAddr := vmctx.req.Unwrap().UTXO().Features().Expiry()
	if expiry == nil {
		return nil
	}

	// TODO maybe iota.go has a function check(output/ expiry time data, own time assumptions) --> true/false
	// To check is output is unlockable based on time assumptions
	// Better reuse logic

	windowFrom := vmctx.finalStateTimestamp.Add(-ExpiryUnlockSafetyWindowDuration)
	windowTo := vmctx.finalStateTimestamp.Add(ExpiryUnlockSafetyWindowDuration)
	if expiry.Time.After(windowFrom) && expiry.Time.Before(windowTo) {
		return xerrors.Errorf("can't be consumed in the expire safety window close to v", expiry.Time)
	}
	milestoneFrom := vmctx.task.TimeAssumption.MilestoneIndex - ExpiryUnlockSafetyWindowMilestone
	milestoneTo := vmctx.task.TimeAssumption.MilestoneIndex + ExpiryUnlockSafetyWindowMilestone
	if milestoneFrom <= expiry.MilestoneIndex && expiry.MilestoneIndex <= milestoneTo {
		return xerrors.Errorf("can't be consumed in the expire safety window between milestones #%d and #%d",
			milestoneFrom, milestoneTo)
	}
	// it is not in the safety window, so it can be consumed either by the chain or by the somebody else
	if vmctx.finalStateTimestamp.After(expiry.Time) {
		if senderAddr.Equal(vmctx.task.AnchorOutput.AliasID.ToAddress()) {
			// this chain is a sender
			// the request came back after expiration
			// TODO ignoring temporary. Must be processed as returned request by consuming it back
			return xerrors.Errorf("came back expired after %v. Ignore for now", expiry.Time)
		} else {
			// somebody else is a sender
			return xerrors.Errorf("expired after %v", expiry.Time)
		}
	}
	return nil
}

// checkReasonReturnAmount skipping anything with return amounts in this version. There's no risk to lose funds
func (vmctx *VMContext) checkReasonReturnAmount() error {
	if _, ok := vmctx.req.Unwrap().UTXO().Features().ReturnAmount(); ok {
		return xerrors.Errorf("return amount feature not supported in this version")
	}
	return nil
}
