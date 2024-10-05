package commondocumentation

import docvalues "config-lsp/doc-values"

var JfsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Character set to use for converting from Unicode to ASCII. The default is to do no conversion. Use iocharset=utf8 for UTF8 translations. This requires CONFIG_NLS_UTF8 to be set in the kernel .config file.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
	docvalues.CreateEnumStringWithDoc(
		"resize",
		"Resize the volume to value blocks. JFS only supports growing a volume, not shrinking it. This option is only valid during a remount, when the volume is mounted read-write. The resize keyword with no value will grow the volume to the full size of the partition.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"errors",
		"Define the behavior when an error is encountered. (Either ignore errors and just mark the filesystem erroneous and continue, or remount the filesystem read-only, or panic and halt the system.)",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("continue"),
			docvalues.CreateEnumString("remount-ro"),
			docvalues.CreateEnumString("panic"),
		},
	},
}

var JfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"nointegrity",
		"Do not write to the journal. The primary use of this option is to allow for higher performance when restoring a volume from backup media. The integrity of the volume is not guaranteed if the system abnormally ends.",
	),
	docvalues.CreateEnumStringWithDoc(
		"integrity",
		"Default. Commit metadata changes to the journal. Use this option to remount a volume where the nointegrity option was previously specified in order to restore normal behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noquota",
		"This option is accepted but ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"quota",
		"This option is accepted but ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"usrquota",
		"This option is accepted but ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"grpquota",
		"This option is accepted but ignored.",
	),
}
