// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "../wasmlib"

export const ScName  = "fairroulette";
export const HScName = new wasmlib.ScHname(0xdf79d138);

export const ParamNumber     = "number";
export const ParamPlayPeriod = "playPeriod";

export const ResultLastWinningNumber = "lastWinningNumber";
export const ResultRoundNumber       = "roundNumber";
export const ResultRoundStartedAt    = "roundStartedAt";
export const ResultRoundStatus       = "roundStatus";

export const StateBets              = "bets";
export const StateLastWinningNumber = "lastWinningNumber";
export const StatePlayPeriod        = "playPeriod";
export const StateRoundNumber       = "roundNumber";
export const StateRoundStartedAt    = "roundStartedAt";
export const StateRoundStatus       = "roundStatus";

export const FuncForcePayout       = "forcePayout";
export const FuncForceReset        = "forceReset";
export const FuncPayWinners        = "payWinners";
export const FuncPlaceBet          = "placeBet";
export const FuncPlayPeriod        = "playPeriod";
export const ViewLastWinningNumber = "lastWinningNumber";
export const ViewRoundNumber       = "roundNumber";
export const ViewRoundStartedAt    = "roundStartedAt";
export const ViewRoundStatus       = "roundStatus";

export const HFuncForcePayout       = new wasmlib.ScHname(0x555a4c4f);
export const HFuncForceReset        = new wasmlib.ScHname(0xa331951e);
export const HFuncPayWinners        = new wasmlib.ScHname(0xfb2b0144);
export const HFuncPlaceBet          = new wasmlib.ScHname(0xdfba7d1b);
export const HFuncPlayPeriod        = new wasmlib.ScHname(0xcb94b293);
export const HViewLastWinningNumber = new wasmlib.ScHname(0x2f5f09fe);
export const HViewRoundNumber       = new wasmlib.ScHname(0x0dcfe520);
export const HViewRoundStartedAt    = new wasmlib.ScHname(0x725de8b4);
export const HViewRoundStatus       = new wasmlib.ScHname(0x145053b5);
