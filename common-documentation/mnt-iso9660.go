package commondocumentation

import docvalues "config-lsp/doc-values"

var Iso9660DocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"check",
		"With check=relaxed, a filename is first converted to lower case before doing the lookup. This is probably only meaningful together with norock and map=normal. (Default: check=strict.)",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("relaxed"),
			docvalues.CreateEnumString("strict"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Give all files in the filesystem the indicated user id, possibly overriding the information found in the Rock Ridge extensions. (Default: uid=0.)",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Give all files in the filesystem the indicated group id, possibly overriding the information found in the Rock Ridge extensions. (Default: gid=0.)",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"map",
		"For non-Rock Ridge volumes, normal name translation maps upper to lower case ASCII, drops a trailing ';1', and converts ';' to '.'. With map=off no name translation is done. See norock. (Default: map=normal.) map=acorn is like map=normal but also apply Acorn extensions if present.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("normal"),
			docvalues.CreateEnumString("off"),
			docvalues.CreateEnumString("acorn"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"For non-Rock Ridge volumes, give all files the indicated mode. (Default: read and execute permission for everybody.) Octal mode values require a leading 0.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"block",
		"Set the block size to the indicated value. (Default: block=1024.)",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("512"),
			docvalues.CreateEnumString("1024"),
			docvalues.CreateEnumString("2048"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"session",
		"Select number of session on a multisession CD.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"sbsector",
		"Session begins from sector xxx.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Character set to use for converting 16 bit Unicode characters on CD to 8 bit characters. The default is iso8859-1.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
	docvalues.CreateEnumStringWithDoc(
		"conv",
		"This option is obsolete and may fail or being ignored.",
	): docvalues.StringValue{},
}

var Iso9660DocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"norock",
		"Disable the use of Rock Ridge extensions, even if available. Cf. map.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nojoliet",
		"Disable the use of Microsoft Joliet extensions, even if available. Cf. map.",
	),
	docvalues.CreateEnumStringWithDoc(
		"unhide",
		"Also show hidden and associated files. (If the ordinary files and the associated or hidden files have the same filenames, this may make the ordinary files inaccessible.)",
	),
	docvalues.CreateEnumStringWithDoc(
		"cruft",
		"If the high byte of the file length contains other garbage, set this mount option to ignore the high order bits of the file length. This implies that a file cannot be larger than 16 MB.",
	),
	docvalues.CreateEnumStringWithDoc(
		"utf8",
		"Convert 16 bit Unicode characters on CD to UTF-8.",
	),
}
