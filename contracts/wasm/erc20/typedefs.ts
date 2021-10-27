// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"

export { MapAgentIDToImmutableInt64 as ImmutableAllowancesForAgent };

export class MapAgentIDToImmutableInt64 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getInt64(key: wasmlib.ScAgentID): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.objID, key.getKeyID());
    }
}

export { MapAgentIDToMutableInt64 as MutableAllowancesForAgent };

export class MapAgentIDToMutableInt64 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID)
    }

    getInt64(key: wasmlib.ScAgentID): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.objID, key.getKeyID());
    }
}
