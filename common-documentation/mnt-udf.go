package commondocumentation

import docvalues "config-lsp/doc-values"

var UdfDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Make all files in the filesystem belong to the given user. uid=forget can be specified independently of (or usually in addition to) uid=<user> and results in UDF not storing uids to the media. In fact the recorded uid is the 32-bit overflow uid -1 as defined by the UDF standard. The value is given as either <user> which is a valid user name or the corresponding decimal user id, or the special string 'forget'.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Make all files in the filesystem belong to the given group. gid=forget can be specified independently of (or usually in addition to) gid=<group> and results in UDF not storing gids to the media. In fact the recorded gid is the 32-bit overflow gid -1 as defined by the UDF standard. The value is given as either <group> which is a valid group name or the corresponding decimal group id, or the special string 'forget'.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"umask",
		"Mask out the given permissions from all inodes read from the filesystem. The value is given in octal.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"If mode= is set the permissions of all non-directory inodes read from the filesystem will be set to the given mode. The value is given in octal.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"dmode",
		"If dmode= is set the permissions of all directory inodes read from the filesystem will be set to the given dmode. The value is given in octal.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"bs",
		"Set the block size. Default value prior to kernel version 2.6.30 was 2048. Since 2.6.30 and prior to 4.11 it was logical device block size with fallback to 2048. Since 4.11 it is logical block size with fallback to any valid block size between logical device block size and 4096.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Set the NLS character set. This requires kernel compiled with CONFIG_UDF_NLS option.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
	docvalues.CreateEnumStringWithDoc(
		"session",
		"Select the session number for multi-session recorded optical media. (default= last session)",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"anchor",
		"Override standard anchor location. (default= 256)",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"lastblock",
		"Set the last block of the filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"conv",
		"This option is obsolete and may fail or being ignored.",
	): docvalues.StringValue{},
}

var UdfDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"unhide",
		"Show otherwise hidden files.",
	),
	docvalues.CreateEnumStringWithDoc(
		"undelete",
		"Show deleted files in lists.",
	),
	docvalues.CreateEnumStringWithDoc(
		"adinicb",
		"Embed data in the inode. (default)",
	),
	docvalues.CreateEnumStringWithDoc(
		"noadinicb",
		"Don’t embed data in the inode.",
	),
	docvalues.CreateEnumStringWithDoc(
		"shortad",
		"Use short UDF address descriptors.",
	),
	docvalues.CreateEnumStringWithDoc(
		"longad",
		"Use long UDF address descriptors. (default)",
	),
	docvalues.CreateEnumStringWithDoc(
		"nostrict",
		"Unset strict conformance.",
	),
	docvalues.CreateEnumStringWithDoc(
		"utf8",
		"Set the UTF-8 character set.",
	),
	docvalues.CreateEnumStringWithDoc(
		"novrs",
		"Ignore the Volume Recognition Sequence and attempt to mount anyway.",
	),
	docvalues.CreateEnumStringWithDoc(
		"session",
		"Select the session number for multi-session recorded optical media. (default= last session)",
	),
	docvalues.CreateEnumStringWithDoc(
		"anchor",
		"Override standard anchor location. (default= 256)",
	),
	docvalues.CreateEnumStringWithDoc(
		"lastblock",
		"Set the last block of the filesystem.",
	),
}
