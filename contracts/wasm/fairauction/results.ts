// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"
import * as sc from "./index";

export class ImmutableGetInfoResults extends wasmlib.ScMapID {

    bidders(): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.mapID, sc.idxMap[sc.IdxResultBidders]);
    }

    color(): wasmlib.ScImmutableColor {
        return new wasmlib.ScImmutableColor(this.mapID, sc.idxMap[sc.IdxResultColor]);
    }

    creator(): wasmlib.ScImmutableAgentID {
        return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxResultCreator]);
    }

    deposit(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultDeposit]);
    }

    description(): wasmlib.ScImmutableString {
        return new wasmlib.ScImmutableString(this.mapID, sc.idxMap[sc.IdxResultDescription]);
    }

    duration(): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.mapID, sc.idxMap[sc.IdxResultDuration]);
    }

    highestBid(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultHighestBid]);
    }

    highestBidder(): wasmlib.ScImmutableAgentID {
        return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxResultHighestBidder]);
    }

    minimumBid(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultMinimumBid]);
    }

    numTokens(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultNumTokens]);
    }

    ownerMargin(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultOwnerMargin]);
    }

    whenStarted(): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultWhenStarted]);
    }
}

export class MutableGetInfoResults extends wasmlib.ScMapID {

    bidders(): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.mapID, sc.idxMap[sc.IdxResultBidders]);
    }

    color(): wasmlib.ScMutableColor {
        return new wasmlib.ScMutableColor(this.mapID, sc.idxMap[sc.IdxResultColor]);
    }

    creator(): wasmlib.ScMutableAgentID {
        return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxResultCreator]);
    }

    deposit(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultDeposit]);
    }

    description(): wasmlib.ScMutableString {
        return new wasmlib.ScMutableString(this.mapID, sc.idxMap[sc.IdxResultDescription]);
    }

    duration(): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.mapID, sc.idxMap[sc.IdxResultDuration]);
    }

    highestBid(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultHighestBid]);
    }

    highestBidder(): wasmlib.ScMutableAgentID {
        return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxResultHighestBidder]);
    }

    minimumBid(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultMinimumBid]);
    }

    numTokens(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultNumTokens]);
    }

    ownerMargin(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultOwnerMargin]);
    }

    whenStarted(): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultWhenStarted]);
    }
}
