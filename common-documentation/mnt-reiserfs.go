package commondocumentation

import docvalues "config-lsp/doc-values"

var ReiserfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"conv",
		"Instructs version 3.6 reiserfs software to mount a version 3.5 filesystem, using the 3.6 format for newly created objects. This filesystem will no longer be compatible with reiserfs 3.5 tools.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"hash",
		"Choose which hash function reiserfs will use to find files within directories.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("rupasov"),
			docvalues.CreateEnumString("tea"),
			docvalues.CreateEnumString("r5"),
			docvalues.CreateEnumString("detect"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"hashed_relocation",
		"Tunes the block allocator. This may provide performance improvements in some situations.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"no_unhashed_relocation",
		"Tunes the block allocator. This may provide performance improvements in some situations.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"noborder",
		"Disable the border allocator algorithm invented by Yury Yu. Rupasov. This may provide performance improvements in some situations.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"nolog",
		"Disable journaling. This will provide slight performance improvements in some situations at the cost of losing reiserfs’s fast recovery from crashes. Even with this option turned on, reiserfs still performs all journaling operations, save for actual writes into its journaling area. Implementation of nolog is a work in progress.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"notail",
		"By default, reiserfs stores small files and 'file tails' directly into its tree. This confuses some utilities such as lilo(8). This option is used to disable packing of files into the tree.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"replayonly",
		"Replay the transactions which are in the journal, but do not actually mount the filesystem. Mainly used by reiserfsck.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"resize",
		"A remount option which permits online expansion of reiserfs partitions. Instructs reiserfs to assume that the device has number blocks. This option is designed for use with devices which are under logical volume management (LVM). There is a special resizer utility which can be obtained from ftp://ftp.namesys.com/pub/reiserfsprogs.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"user_xattr",
		"Enable Extended User Attributes. See the [attr(1)](https://www.man7.org/linux/man-pages/man1/attr.1.html) manual page.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"acl",
		"Enable POSIX Access Control Lists. See the [acl(5)](https://www.man7.org/linux/man-pages/man5/acl.5.html) manual page.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"barrier",
		"This disables / enables the use of write barriers in the journaling code. barrier=none disables, barrier=flush enables (default). This also requires an IO stack which can support barriers, and if reiserfs gets an error on a barrier write, it will disable barriers again with a warning. Write barriers enforce proper on-disk ordering of journal commits, making volatile disk write caches safe to use, at some performance penalty. If your disks are battery-backed in one way or another, disabling barriers may safely improve performance.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("none"),
			docvalues.CreateEnumString("flush"),
		},
	},
}

var ReiserfsDocumentationEnums = []docvalues.EnumString{}
