// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"

export { ArrayOfImmutableString as ImmutableStringArray };

export class ArrayOfImmutableString {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getString(index: i32): wasmlib.ScImmutableString {
        return new wasmlib.ScImmutableString(this.objID, new wasmlib.Key32(index));
    }
}

export { ArrayOfMutableString as MutableStringArray };

export class ArrayOfMutableString {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getString(index: i32): wasmlib.ScMutableString {
        return new wasmlib.ScMutableString(this.objID, new wasmlib.Key32(index));
    }
}
