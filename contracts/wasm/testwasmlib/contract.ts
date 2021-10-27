// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"
import * as sc from "./index";

export class ArrayClearCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncArrayClear);
    params: sc.MutableArrayClearParams = new sc.MutableArrayClearParams();
}

export class ArrayCreateCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncArrayCreate);
    params: sc.MutableArrayCreateParams = new sc.MutableArrayCreateParams();
}

export class ArraySetCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncArraySet);
    params: sc.MutableArraySetParams = new sc.MutableArraySetParams();
}

export class ParamTypesCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncParamTypes);
    params: sc.MutableParamTypesParams = new sc.MutableParamTypesParams();
}

export class ArrayLengthCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewArrayLength);
    params: sc.MutableArrayLengthParams = new sc.MutableArrayLengthParams();
    results: sc.ImmutableArrayLengthResults = new sc.ImmutableArrayLengthResults();
}

export class ArrayValueCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewArrayValue);
    params: sc.MutableArrayValueParams = new sc.MutableArrayValueParams();
    results: sc.ImmutableArrayValueResults = new sc.ImmutableArrayValueResults();
}

export class BlockRecordCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewBlockRecord);
    params: sc.MutableBlockRecordParams = new sc.MutableBlockRecordParams();
    results: sc.ImmutableBlockRecordResults = new sc.ImmutableBlockRecordResults();
}

export class BlockRecordsCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewBlockRecords);
    params: sc.MutableBlockRecordsParams = new sc.MutableBlockRecordsParams();
    results: sc.ImmutableBlockRecordsResults = new sc.ImmutableBlockRecordsResults();
}

export class IotaBalanceCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewIotaBalance);
    results: sc.ImmutableIotaBalanceResults = new sc.ImmutableIotaBalanceResults();
}

export class ScFuncs {

    static arrayClear(ctx: wasmlib.ScFuncCallContext): ArrayClearCall {
        let f = new ArrayClearCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static arrayCreate(ctx: wasmlib.ScFuncCallContext): ArrayCreateCall {
        let f = new ArrayCreateCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static arraySet(ctx: wasmlib.ScFuncCallContext): ArraySetCall {
        let f = new ArraySetCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static paramTypes(ctx: wasmlib.ScFuncCallContext): ParamTypesCall {
        let f = new ParamTypesCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static arrayLength(ctx: wasmlib.ScViewCallContext): ArrayLengthCall {
        let f = new ArrayLengthCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static arrayValue(ctx: wasmlib.ScViewCallContext): ArrayValueCall {
        let f = new ArrayValueCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static blockRecord(ctx: wasmlib.ScViewCallContext): BlockRecordCall {
        let f = new BlockRecordCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static blockRecords(ctx: wasmlib.ScViewCallContext): BlockRecordsCall {
        let f = new BlockRecordsCall();
        f.func.setPtrs(f.params, f.results);
        return f;
    }

    static iotaBalance(ctx: wasmlib.ScViewCallContext): IotaBalanceCall {
        let f = new IotaBalanceCall();
        f.func.setPtrs(null, f.results);
        return f;
    }
}
