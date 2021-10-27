// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"
import * as sc from "./index";

export class ImmutableStoreBlobResults extends wasmlib.ScMapID {

    hash(): wasmlib.ScImmutableHash {
        return new wasmlib.ScImmutableHash(this.mapID, wasmlib.Key32.fromString(sc.ResultHash));
    }
}

export class MutableStoreBlobResults extends wasmlib.ScMapID {

    hash(): wasmlib.ScMutableHash {
        return new wasmlib.ScMutableHash(this.mapID, wasmlib.Key32.fromString(sc.ResultHash));
    }
}

export class ImmutableGetBlobFieldResults extends wasmlib.ScMapID {

    bytes(): wasmlib.ScImmutableBytes {
        return new wasmlib.ScImmutableBytes(this.mapID, wasmlib.Key32.fromString(sc.ResultBytes));
    }
}

export class MutableGetBlobFieldResults extends wasmlib.ScMapID {

    bytes(): wasmlib.ScMutableBytes {
        return new wasmlib.ScMutableBytes(this.mapID, wasmlib.Key32.fromString(sc.ResultBytes));
    }
}

export class MapStringToImmutableInt32 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getInt32(key: string): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.objID, wasmlib.Key32.fromString(key).getKeyID());
    }
}

export class ImmutableGetBlobInfoResults extends wasmlib.ScMapID {

    blobSizes(): sc.MapStringToImmutableInt32 {
        return new sc.MapStringToImmutableInt32(this.mapID);
    }
}

export class MapStringToMutableInt32 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID)
    }

    getInt32(key: string): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.objID, wasmlib.Key32.fromString(key).getKeyID());
    }
}

export class MutableGetBlobInfoResults extends wasmlib.ScMapID {

    blobSizes(): sc.MapStringToMutableInt32 {
        return new sc.MapStringToMutableInt32(this.mapID);
    }
}

export class MapHashToImmutableInt32 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getInt32(key: wasmlib.ScHash): wasmlib.ScImmutableInt32 {
        return new wasmlib.ScImmutableInt32(this.objID, key.getKeyID());
    }
}

export class ImmutableListBlobsResults extends wasmlib.ScMapID {

    blobSizes(): sc.MapHashToImmutableInt32 {
        return new sc.MapHashToImmutableInt32(this.mapID);
    }
}

export class MapHashToMutableInt32 {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID)
    }

    getInt32(key: wasmlib.ScHash): wasmlib.ScMutableInt32 {
        return new wasmlib.ScMutableInt32(this.objID, key.getKeyID());
    }
}

export class MutableListBlobsResults extends wasmlib.ScMapID {

    blobSizes(): sc.MapHashToMutableInt32 {
        return new sc.MapHashToMutableInt32(this.mapID);
    }
}
