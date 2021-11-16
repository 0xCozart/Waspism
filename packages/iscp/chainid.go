// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package iscp

import (
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/mr-tron/base58"

	"github.com/iotaledger/hive.go/marshalutil"

	"github.com/iotaledger/wasp/packages/hashing"
	"golang.org/x/xerrors"
)

const ChainIDLength = iotago.AliasIDLength

// ChainID represents the global identifier of the chain
// It is wrapped AliasAddress, an address without a private key behind
type ChainID iotago.AliasID

// NewChainID creates new chain ID from alias address
func NewChainID(addr iotago.AliasID) ChainID {
	return ChainID(addr)
}

// ChainIDFromAddress creates a chainIDD from alias address. Returns and error if not an alias address type
// Deprecated:
func ChainIDFromAddress(addr iotago.Address) (ChainID, error) {
	if addr.Type() != iotago.AddressAlias {
		return ChainID{}, xerrors.New("chain id must be an alias address")
	}
	return ChainID{}, nil
}

// ChainIDFromBytes reconstructs a ChainID from its binary representation.
func ChainIDFromBytes(data []byte) (ChainID, error) {
	var ret ChainID
	if len(ret) != len(data) {
		return ChainID{}, xerrors.New("ChainIDFromBase58: wrong data length")
	}
	copy(ret[:], data)
	return ret, nil
}

// ChainIDFromBase58 constructor decodes base58 string to the ChainID
func ChainIDFromBase58(b58 string) (ChainID, error) {
	bin, err := base58.Decode(b58)
	if err != nil {
		return ChainID{}, err
	}
	return ChainIDFromBytes(bin)
}

// TODO adjust to iotago style
// ChainIDFromMarshalUtil reads from Marshalutil
func ChainIDFromMarshalUtil(mu *marshalutil.MarshalUtil) (ChainID, error) {
	bin, err := mu.ReadBytes(ChainIDLength)
	if err != nil {
		return ChainID{}, err
	}
	return ChainIDFromBytes(bin)
}

// RandomChainID creates a random chain ID. Used for testing only
func RandomChainID(seed ...[]byte) ChainID {
	var h hashing.HashValue
	if len(seed) > 0 {
		h = hashing.HashData(seed[0])
	} else {
		h = hashing.RandomHash(nil)
	}
	ret, _ := ChainIDFromBytes(h[:ChainIDLength])
	return ret
}

// Equals for using
func (chid *ChainID) Equals(chid1 *ChainID) bool {
	return chid == chid1
}

func (chid *ChainID) Base58() string {
	return base58.Encode(chid[:])
}

// String human readable form (base58 encoding)
func (chid *ChainID) String() string {
	return "$/" + chid.Base58()
}

func (chid *ChainID) AsAddress() iotago.Address {
	ret := iotago.AliasAddress(*chid)
	return &ret
}

func (chid *ChainID) AsAliasAddress() iotago.AliasAddress {
	return iotago.AliasAddress(*chid)
}
