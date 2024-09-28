package commondocumentation

import docvalues "config-lsp/doc-values"

var ReiserfsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"hash",
		"Choose which hash function reiserfs will use to find files within directories.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"rupasov",
				"A hash invented by Yury Yu. Rupasov. It is fast and preserves locality, mapping lexicographically close file names to close hash values. This option should not be used, as it causes a high probability of hash collisions.",
			),
			docvalues.CreateEnumStringWithDoc(
				"tea",
				"A Davis-Meyer function implemented by Jeremy Fitzhardinge. It uses hash permuting bits in the name. It gets high randomness and, therefore, low probability of hash collisions at some CPU cost. This may be used if EHASHCOLLISION errors are experienced with the r5 hash.",
			),
			docvalues.CreateEnumStringWithDoc(
				"r5",
				"A modified version of the rupasov hash. It is used by default and is the best choice unless the filesystem has huge directories and unusual file-name patterns.",
			),
			docvalues.CreateEnumStringWithDoc(
				"detect",
				"Instructs mount to detect which hash function is in use by examining the filesystem being mounted, and to write this information into the reiserfs superblock. This is only useful on the first mount of an old format filesystem.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"resize",
		"A remount option which permits online expansion of reiserfs partitions. Instructs reiserfs to assume that the device has number blocks. This option is designed for use with devices which are under logical volume management (LVM). There is a special resizer utility which can be obtained from ftp://ftp.namesys.com/pub/reiserfsprogs.",
	): docvalues.NumberValue{Min: &zero},
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

var ReiserfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"conv",
		"Instructs version 3.6 reiserfs software to mount a version 3.5 filesystem, using the 3.6 format for newly created objects. This filesystem will no longer be compatible with reiserfs 3.5 tools.",
	),
	docvalues.CreateEnumStringWithDoc(
		"hashed_relocation",
		"Tunes the block allocator. This may provide performance improvements in some situations.",
	),
	docvalues.CreateEnumStringWithDoc(
		"no_unhashed_relocation",
		"Tunes the block allocator. This may provide performance improvements in some situations.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noborder",
		"Disable the border allocator algorithm invented by Yury Yu. Rupasov. This may provide performance improvements in some situations.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nolog",
		"Disable journaling. This will provide slight performance improvements in some situations at the cost of losing reiserfs’s fast recovery from crashes. Even with this option turned on, reiserfs still performs all journaling operations, save for actual writes into its journaling area. Implementation of nolog is a work in progress.",
	),
	docvalues.CreateEnumStringWithDoc(
		"notail",
		"By default, reiserfs stores small files and 'file tails' directly into its tree. This confuses some utilities such as lilo(8). This option is used to disable packing of files into the tree.",
	),
	docvalues.CreateEnumStringWithDoc(
		"replayonly",
		"Replay the transactions which are in the journal, but do not actually mount the filesystem. Mainly used by reiserfsck.",
	),
	docvalues.CreateEnumStringWithDoc(
		"user_xattr",
		"Enable Extended User Attributes. See the [attr(1)](https://www.man7.org/linux/man-pages/man1/attr.1.html) manual page.",
	),
	docvalues.CreateEnumStringWithDoc(
		"acl",
		"Enable POSIX Access Control Lists. See the [acl(5)](https://www.man7.org/linux/man-pages/man5/acl.5.html) manual page.",
	),
}
