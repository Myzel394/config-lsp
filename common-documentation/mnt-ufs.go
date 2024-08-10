package commondocumentation

import docvalues "config-lsp/doc-values"

var UfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"ufstype",
		"UFS is a filesystem widely used in different operating systems. The problem are differences among implementations. Features of some implementations are undocumented, so its hard to recognize the type of ufs automatically. That’s why the user must specify the type of ufs by mount option.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("old"),
			docvalues.CreateEnumString("44bsd"),
			docvalues.CreateEnumString("ufs2"),
			docvalues.CreateEnumString("5xbsd"),
			docvalues.CreateEnumString("sun"),
			docvalues.CreateEnumString("sunx86"),
			docvalues.CreateEnumString("hp"),
			docvalues.CreateEnumString("nextstep"),
			docvalues.CreateEnumString("nextstep-cd"),
			docvalues.CreateEnumString("openstep"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"onerror",
		"Set behavior on error.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("panic"),
			docvalues.CreateEnumString("lock"),
			docvalues.CreateEnumString("umount"),
			docvalues.CreateEnumString("repair"),
		},
	},
}

var UfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"old",
		"Old format of ufs, this is the default, read only. (Don’t forget to give the -r option.)",
	),
	docvalues.CreateEnumStringWithDoc(
		"44bsd",
		"For filesystems created by a BSD-like system (NetBSD, FreeBSD, OpenBSD).",
	),
	docvalues.CreateEnumStringWithDoc(
		"ufs2",
		"Used in FreeBSD 5.x supported as read-write.",
	),
	docvalues.CreateEnumStringWithDoc(
		"5xbsd",
		"Synonym for ufs2.",
	),
	docvalues.CreateEnumStringWithDoc(
		"sun",
		"For filesystems created by SunOS or Solaris on Sparc.",
	),
	docvalues.CreateEnumStringWithDoc(
		"sunx86",
		"For filesystems created by Solaris on x86.",
	),
	docvalues.CreateEnumStringWithDoc(
		"hp",
		"For filesystems created by HP-UX, read-only.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nextstep",
		"For filesystems created by NeXTStep (on NeXT station) (currently read only).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nextstep-cd",
		"For NextStep CDROMs (block_size == 2048), read-only.",
	),
	docvalues.CreateEnumStringWithDoc(
		"openstep",
		"For filesystems created by OpenStep (currently read only). The same filesystem type is also used by macOS.",
	),
	docvalues.CreateEnumStringWithDoc(
		"panic",
		"If an error is encountered, cause a kernel panic.",
	),
	docvalues.CreateEnumStringWithDoc(
		"lock",
		"These mount options don’t do anything at present; when an error is encountered only a console message is printed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"umount",
		"These mount options don’t do anything at present; when an error is encountered only a console message is printed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"repair",
		"These mount options don’t do anything at present; when an error is encountered only a console message is printed.",
	),
}
