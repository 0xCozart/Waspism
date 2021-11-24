// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package iscp

import (
	"math/big"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"

	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
)

// SandboxBase is the common interface of Sandbox and SandboxView
type SandboxBase interface {
	Balance
	// AccountID returns the agentID of the current contract
	AccountID() *AgentID
	// Params returns the parameters of the current call
	Params() dict.Dict
	// ChainID returns the chain ID
	ChainID() *ChainID
	// ChainOwnerID returns the AgentID of the current owner of the chain
	ChainOwnerID() *AgentID
	// Contract returns the Hname of the contract in the current chain
	Contract() Hname
	// ContractCreator returns the agentID that deployed the contract
	ContractCreator() *AgentID
	// Timestamp returns the UnixNano timestamp of the current state
	Timestamp() int64
	// Log returns a logger that outputs on the local machine. It includes Panicf method
	Log() LogInterface
	// Utils provides access to common necessary functionality
	Utils() Utils
	// Gas returns interface for gas related functions
	Gas() Gas
}

type Balance interface {
	// BalanceIotas returns number of iotas in the balance of the smart contract
	BalanceIotas() uint64
	// BalanceNativeToken returns number of native token or nil if it is empty
	BalanceNativeToken(id *iotago.NativeTokenID) *big.Int
	// Assets returns all assets: iotas and native tokens
	Assets() *Assets
}

// Sandbox is an interface given to the processor to access the VMContext
// and virtual state, transaction builder and request parameters through it.
type Sandbox interface {
	SandboxBase

	// State k/v store of the current call (in the context of the smart contract)
	State() kv.KVStore
	// Request return the request in the context of which the smart contract is called
	Request() Request

	// Call calls the entry point of the contract with parameters and Transfer.
	// If the entry point is full entry point, Transfer tokens are moved between caller's and
	// target contract's accounts (if enough). If the entry point is view, 'Transfer' has no effect
	Call(target, entryPoint Hname, params dict.Dict, transfer *Assets) (dict.Dict, error)
	// Caller is the agentID of the caller.
	Caller() *AgentID
	// DeployContract deploys contract on the same chain. 'initParams' are passed to the 'init' entry point
	DeployContract(programHash hashing.HashValue, name string, description string, initParams dict.Dict) error
	// Event publishes "vmmsg" message through Publisher on nanomsg. It also logs locally, but it is not the same thing
	Event(msg string)
	// GetEntropy 32 random bytes based on the hash of the current state transaction
	GetEntropy() hashing.HashValue // 32 bytes of deterministic and unpredictably random data
	// IncomingTransfer return colored balances transferred by the call. They are already accounted into the Balances()
	IncomingTransfer() *Assets
	// Send one generic method for sending assets with ledgerstate.ExtendedLockedOutput
	// replaces TransferToAddress and Post1Request
	Send(target iotago.Address, assets *Assets, metadata *SendMetadata, options ...*SendOptions)
	// BlockContext Internal for use in native hardcoded contracts
	BlockContext(construct func(sandbox Sandbox) interface{}, onClose func(interface{})) interface{}
	// StateAnchor properties of the anchor output
	StateAnchor() StateAnchor
}

type Gas interface {
	Burn(uint64)
	Budget() uint64
}

// properties of the anchor output/transaction in the current context
type StateAnchor interface {
	StateController() iotago.Address
	GovernanceController() iotago.Address
	StateIndex() uint32
	OutputID() iotago.UTXOInput
	StateData() StateData
}

type SendOptions struct { // TODO
}

// RequestMetadata represents content of the data payload of the output
type SendMetadata struct {
	TargetContract Hname
	EntryPoint     Hname
	Params         dict.Dict
	Transfer       *Assets
	GasBudget      uint64
}

// Utils implement various utilities which are faster on host side than on wasm VM
// Implement deterministic stateless computations
type Utils interface {
	Base58() Base58
	Hashing() Hashing
	ED25519() ED25519
	BLS() BLS
}

type Hashing interface {
	Blake2b(data []byte) hashing.HashValue
	Sha3(data []byte) hashing.HashValue
	Hname(name string) Hname
}

type Base58 interface {
	Decode(s string) ([]byte, error)
	Encode(data []byte) string
}

type ED25519 interface {
	ValidSignature(data []byte, pubKey []byte, signature []byte) bool
	AddressFromPublicKey(pubKey []byte) (ledgerstate.Address, error)
}

type BLS interface {
	ValidSignature(data []byte, pubKey []byte, signature []byte) bool
	AddressFromPublicKey(pubKey []byte) (ledgerstate.Address, error)
	AggregateBLSSignatures(pubKeysBin [][]byte, sigsBin [][]byte) ([]byte, []byte, error)
}
