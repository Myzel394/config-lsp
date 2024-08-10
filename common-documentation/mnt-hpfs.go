package commondocumentation

import docvalues "config-lsp/doc-values"

var HpfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the owner and group of all files. (Default: the UID and GID of the current process.)",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the owner and group of all files. (Default: the UID and GID of the current process.)",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"umask",
		"Set the umask (the bitmask of the permissions that are not present). The default is the umask of the current process. The value is given in octal.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"case",
		"Convert all files names to lower case, or leave them. (Default: case=lower.)",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("lower"),
			docvalues.CreateEnumString("asis"),
		},
	},
	// Todo: Show error in analyzer
	docvalues.CreateEnumStringWithDoc(
		"conv",
		"This option is obsolete and may fail or being ignored.",
	): docvalues.StringValue{},
}

var HpfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"nocheck",
		"Do not abort mounting when certain consistency checks fail.",
	),
}
