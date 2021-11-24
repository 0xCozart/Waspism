package rstemplates

var RsTemplates = []map[string]string{
	rsCommon,
	cargoToml,
	constsRs,
	contractRs,
	eventsRs,
	funcsRs,
	keysRs,
	libRs,
	modRs,
	paramsRs,
	proxyRs,
	resultsRs,
	stateRs,
	structsRs,
	typedefsRs,
}

var rsCommon = map[string]string{
	// *******************************
	"initGlobals": `
$#set arrayTypeID TYPE_ARRAY
$#set crate 
$#if core setArrayTypeID
`,
	// *******************************
	"setArrayTypeID": `
$#set arrayTypeID TYPE_ARRAY16
$#set crate (crate)
`,
	// *******************************
	"rsHeader": `
$#if core useCrate useWasmLib
`,
	// *******************************
	"modEvents": `
mod events;
`,
	// *******************************
	"modParams": `
mod params;
`,
	// *******************************
	"modResults": `
mod results;
`,
	// *******************************
	"modStructs": `
mod structs;
`,
	// *******************************
	"modTypeDefs": `
mod typedefs;
`,
	// *******************************
	"useCrate": `
use crate::*;
`,
	// *******************************
	"useCoreContract": `
use crate::$package::*;
`,
	// *******************************
	"useEvents": `
use crate::events::*;
`,
	// *******************************
	"useHost": `
use crate::host::*;
`,
	// *******************************
	"useParams": `
use crate::params::*;
`,
	// *******************************
	"useResults": `
use crate::results::*;
`,
	// *******************************
	"useStructs": `
use crate::structs::*;
`,
	// *******************************
	"useTypeDefs": `
use crate::typedefs::*;
`,
	// *******************************
	"useWasmLib": `
use wasmlib::*;
`,
	// *******************************
	"setVarType": `
$#set varType TYPE_MAP
$#if array setVarTypeArray
`,
	// *******************************
	"setVarTypeArray": `
$#set varType $arrayTypeID | $fldTypeID
`,
}
