package vmcontext

import "github.com/iotaledger/wasp/packages/iscp/coreutil"

func (vmctx *VMContext) gasSetBudget(gasBudget uint64) {
	vmctx.gasBudget = gasBudget
	vmctx.gasBurned = 0
}

func (vmctx *VMContext) GasBurn(gas uint64) {
	vmctx.gasBurned += gas
	if vmctx.gasBurned > vmctx.gasBudget {
		panic(coreutil.ErrorGasBudgetExceeded)
	}
}

func (vmctx *VMContext) GasBudgetLeft() uint64 {
	if vmctx.gasBudget < vmctx.gasBurned {
		return 0
	}
	return vmctx.gasBudget - vmctx.gasBurned
}

func (vmctx *VMContext) GasBurned() uint64 {
	return vmctx.gasBurned
}
