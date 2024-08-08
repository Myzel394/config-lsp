package commondocumentation

import docvalues "config-lsp/doc-values"

var HfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"creator",
		"Set the creator/type values as shown by the MacOS finder used for creating new files. Default values: '????'.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"type",
		"Set the creator/type values as shown by the MacOS finder used for creating new files. Default values: '????'.",
	): docvalues.StringValue{},
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
		"dir_umask",
		"Set the umask used for all directories. Defaults to the umask of the current process.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"file_umask",
		"Set the umask used for all regular files. Defaults to the umask of the current process.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"umask",
		"Set the umask used for all files and directories. Defaults to the umask of the current process.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"session",
		"Select the CDROM session to mount. Defaults to leaving that decision to the CDROM driver. This option will fail with anything but a CDROM as underlying device.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"part",
		"Select partition number n from the device. Only makes sense for CDROMs. Defaults to not parsing the partition table at all.",
	): docvalues.NumberValue{Min: &zero},
}

var HfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"quiet",
		"Don't complain about invalid mount options.",
	),
}
