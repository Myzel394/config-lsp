package commondocumentation

import docvalues "config-lsp/doc-values"

var DebugfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the owner of the mountpoint.",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the group of the mountpoint.",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"Sets the mode of the mountpoint.",
	): docvalues.MaskModeValue{},
}

var DebugfsDocumentationEnums = []docvalues.EnumString{}
