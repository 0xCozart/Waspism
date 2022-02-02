// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::wasmtypes::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub const SC_HNAME_LENGTH: usize = 4;

#[derive(PartialEq, Clone)]
pub struct ScHname(pub u32);

impl ScHname {
    pub fn new(buf: &[u8]) -> ScHname {
        ScHname(uint32_from_bytes(buf))
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        hname_to_bytes(self)
    }

    pub fn to_string(&self) -> String {
        hname_to_string(self)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub fn hname_decode(dec: &mut WasmDecoder) -> ScHname {
    hname_from_bytes(&dec.fixed_bytes(SC_HNAME_LENGTH))
}

pub fn hname_encode(enc: &mut WasmEncoder, value: &ScHname)  {
    enc.fixed_bytes(&hname_to_bytes(value), SC_HNAME_LENGTH);
}

pub fn hname_from_bytes(buf: &[u8]) -> ScHname {
    if buf.len() == 0 {
        return ScHname(0);
    }
    ScHname(u32::from_le_bytes(buf.try_into().expect("invalid Hname length")))
}

pub fn hname_to_bytes(value: &ScHname) -> Vec<u8> {
    value.0.to_le_bytes().to_vec()
}

pub fn hname_to_string(value: &ScHname) -> String {
    let hexa = "0123456789abcdef".as_bytes();
    let mut res = [0u8; 8];
    let mut val = value.0;
    for n in 0..8 {
        res[7 - n] = hexa[val as usize & 0x0f];
        val >>= 4;
    }
    String::from_utf8(res.to_vec()).expect("WTF? invalid?")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableHname<'a> {
    proxy: Proxy<'a>,
}

impl ScImmutableHname<'_> {
    pub fn new(proxy: Proxy) -> ScImmutableHname {
        ScImmutableHname { proxy }
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn to_string(&self) -> String {
        hname_to_string(&self.value())
    }

    pub fn value(&self) -> ScHname {
        hname_from_bytes(&self.proxy.get())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for mutable ScHname in host container
pub struct ScMutableHname<'a> {
    proxy: Proxy<'a>,
}

impl ScMutableHname<'_> {
    pub fn new(proxy: Proxy) -> ScMutableHname {
        ScMutableHname { proxy }
    }

    pub fn delete(&mut self)  {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&mut self, value: &ScHname) {
        self.proxy.set(&hname_to_bytes(&value));
    }

    pub fn to_string(&self) -> String {
        hname_to_string(&self.value())
    }

    pub fn value(&self) -> ScHname {
        hname_from_bytes(&self.proxy.get())
    }
}
