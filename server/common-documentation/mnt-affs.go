package commondocumentation

import docvalues "config-lsp/doc-values"

var prefixMaxLength = uint32(30)

var AffsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the owner and group of the root of the filesystem (default: uid=gid=0, but with option uid or gid without specified value, the UID and GID of the current process are taken).",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the owner and group of the root of the filesystem (default: uid=gid=0, but with option uid or gid without specified value, the UID and GID of the current process are taken).",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"setuid",
		"Set the owner of all files.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"setgid",
		"Set the group of all files.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"Set the mode of all files to value & 0777 disregarding the original permissions. Add search permission to directories that have read permission. The value is given in octal.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"prefix",
		"Prefix used before volume name, when following a link.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"volume",
		"Prefix (of length at most 30) used before '/' when following a symbolic link.",
	): docvalues.StringValue{MaxLength: &prefixMaxLength},
	docvalues.CreateEnumStringWithDoc(
		"reserved",
		"(Default: 2.) Number of unused blocks at the start of the device.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"root",
		"Give explicitly the location of the root block.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		"bs",
		"Give blocksize. Allowed values are 512, 1024, 2048, 4096.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("512"),
			docvalues.CreateEnumString("1024"),
			docvalues.CreateEnumString("2048"),
			docvalues.CreateEnumString("4096"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"grpquota",
		"These options are accepted but ignored. (However, quota utilities may react to such strings in /etc/fstab.)",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"noquota",
		"These options are accepted but ignored. (However, quota utilities may react to such strings in /etc/fstab.)",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"quota",
		"These options are accepted but ignored. (However, quota utilities may react to such strings in /etc/fstab.)",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"usrquota",
		"These options are accepted but ignored. (However, quota utilities may react to such strings in /etc/fstab.)",
	): docvalues.StringValue{},
}

var AffsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"protect",
		"Do not allow any changes to the protection bits on the filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"usemp",
		"Set UID and GID of the root of the filesystem to the UID and GID of the mount point upon the first sync or umount, and then clear this option. Strange...",
	),
	docvalues.CreateEnumStringWithDoc(
		"verbose",
		"Print an informational message for each successful mount.",
	),
}
