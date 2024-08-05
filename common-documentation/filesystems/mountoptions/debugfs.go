package commondocumentation

import docvalues "config-lsp/doc-values"

var DebugfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the owner of the mountpoint.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the group of the mountpoint.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"Sets the mode of the mountpoint.",
	): docvalues.StringValue{},
}
