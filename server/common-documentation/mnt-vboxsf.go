package commondocumentation

import docvalues "config-lsp/doc-values"

var VboxsfDocumentationAssignable = docvalues.MergeKeyEnumAssignmentMaps(FatDocumentationAssignable, map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"This option sets the character set used for I/O operations. Note that on Linux guests, if the iocharset option is not specified, then the Guest Additions driver will attempt to use the character set specified by the CONFIG_NLS_DEFAULT kernel option. If this option is not set either, then UTF-8 is used.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
	docvalues.CreateEnumStringWithDoc(
		"convertcp",
		"This option specifies the character set used for the shared folder name. This is UTF-8 by default.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
})

var VboxsfDocumentationEnums = []docvalues.EnumString{}
