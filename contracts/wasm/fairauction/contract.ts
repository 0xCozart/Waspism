// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"
import * as sc from "./index";

export class FinalizeAuctionCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncFinalizeAuction);
    params: sc.MutableFinalizeAuctionParams = new sc.MutableFinalizeAuctionParams();
}

export class PlaceBidCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPlaceBid);
    params: sc.MutablePlaceBidParams = new sc.MutablePlaceBidParams();
}

export class SetOwnerMarginCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncSetOwnerMargin);
    params: sc.MutableSetOwnerMarginParams = new sc.MutableSetOwnerMarginParams();
}

export class StartAuctionCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncStartAuction);
    params: sc.MutableStartAuctionParams = new sc.MutableStartAuctionParams();
}

export class GetInfoCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetInfo);
    params: sc.MutableGetInfoParams = new sc.MutableGetInfoParams();
    results: sc.ImmutableGetInfoResults = new sc.ImmutableGetInfoResults();
}

export class ScFuncs {

    static finalizeAuction(ctx: wasmlib.ScFuncCallContext): FinalizeAuctionCall {
        let f = new FinalizeAuctionCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static placeBid(ctx: wasmlib.ScFuncCallContext): PlaceBidCall {
        let f = new PlaceBidCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static setOwnerMargin(ctx: wasmlib.ScFuncCallContext): SetOwnerMarginCall {
        let f = new SetOwnerMarginCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static startAuction(ctx: wasmlib.ScFuncCallContext): StartAuctionCall {
        let f = new StartAuctionCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static getInfo(ctx: wasmlib.ScViewCallContext): GetInfoCall {
        let f = new GetInfoCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }
}
