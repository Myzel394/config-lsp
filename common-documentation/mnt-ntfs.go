package commondocumentation

import docvalues "config-lsp/doc-values"

var NtfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Character set to use when returning file names. Unlike VFAT, NTFS suppresses names that contain nonconvertible characters. Deprecated.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"nls",
		"New name for the option earlier called iocharset.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"utf8",
		"Use UTF-8 for converting file names.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"uni_xlate",
		"For 0 (or 'no' or 'false'), do not use escape sequences for unknown Unicode characters. For 1 (or 'yes' or 'true') or 2, use vfat-style 4-byte escape sequences starting with ':'. Here 2 gives a little-endian encoding and 1 a byteswapped bigendian encoding.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("0"),
			docvalues.CreateEnumString("1"),
			docvalues.CreateEnumString("2"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"posix",
		"If enabled (posix=1), the filesystem distinguishes between upper and lower case. The 8.3 alias names are presented as hard links instead of being suppressed. This option is obsolete.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("0"),
			docvalues.CreateEnumString("1"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the file permission on the filesystem. The umask value is given in octal. By default, the files are owned by root and not readable by somebody else.",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the file permission on the filesystem. The umask value is given in octal. By default, the files are owned by root and not readable by somebody else.",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"umask",
		"Set the file permission on the filesystem. The umask value is given in octal. By default, the files are owned by root and not readable by somebody else.",
	): docvalues.UmaskValue{},
}

var NtfsDocumentationEnums = []docvalues.EnumString{}
