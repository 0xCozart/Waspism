// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::convert::TryInto;

use crate::*;
use crate::wasmtypes::*;

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub const SC_ADDRESS_LENGTH: usize = 33;

#[derive(PartialEq, Clone)]
pub struct ScAddress {
    id: [u8; SC_ADDRESS_LENGTH],
}

impl ScAddress {
    pub fn new(buf: &[u8]) -> ScAddress {
        address_from_bytes(buf)
    }

    pub fn as_agent_id(&self) -> ScAgentID {
        ScAgentID::new(self, &ScHname(0))
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        address_to_bytes(self)
    }

    pub fn to_string(&self) -> String {
        address_to_string(self)
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub fn address_decode(dec: &mut WasmDecoder) -> ScAddress {
    address_from_bytes_unchecked(&dec.fixed_bytes(SC_ADDRESS_LENGTH))
}

pub fn address_encode(enc: &mut WasmEncoder, value: &ScAddress)  {
    enc.fixed_bytes(&value.to_bytes(), SC_ADDRESS_LENGTH);
}

pub fn address_from_bytes(buf: &[u8]) -> ScAddress {
    ScAddress { id: buf.try_into().expect("invalid Address length") }
}

pub fn address_to_bytes(value: &ScAddress) -> Vec<u8> {
    value.id.to_vec()
}

pub fn address_to_string(value: &ScAddress) -> String {
    // TODO standardize human readable string
    base58_encode(&value.id)
}

fn address_from_bytes_unchecked(buf: &[u8]) -> ScAddress {
    ScAddress { id: buf.try_into().expect("invalid Address length") }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

pub struct ScImmutableAddress<'a> {
    proxy: Proxy<'a>,
}

impl ScImmutableAddress<'_> {
    pub fn new(proxy: Proxy) -> ScImmutableAddress {
        ScImmutableAddress { proxy }
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn to_string(&self) -> String {
        address_to_string(&self.value())
    }

    pub fn value(&self) -> ScAddress {
        address_from_bytes(&self.proxy.get())
    }
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

// value proxy for mutable ScAddress in host container
pub struct ScMutableAddress<'a> {
    proxy: Proxy<'a>,
}

impl ScMutableAddress<'_> {
    pub fn new(proxy: Proxy) -> ScMutableAddress {
        ScMutableAddress { proxy }
    }

    pub fn delete(&mut self)  {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&mut self, value: &ScAddress) {
        self.proxy.set(&address_to_bytes(&value));
    }

    pub fn to_string(&self) -> String {
        address_to_string(&self.value())
    }

    pub fn value(&self) -> ScAddress {
        address_from_bytes(&self.proxy.get())
    }
}
