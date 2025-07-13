package commondocumentation

import docvalues "config-lsp/doc-values"

// TODO: Add user properties

var numericProp = docvalues.DataAmountValue{
	AllowedUnits: map[rune]struct{}{
		'k': {},
		'm': {},
		'g': {},
		't': {},
		'p': {},
		'e': {},
		'z': {},
	},
	AllowByteSuffix: true,
	AllowDecimal:    true,
	Base:            docvalues.DataAmountValueBase1024,
}

var ZfsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"context",
		"Sets the SELinux context for all files in the file system. This option is used to set the security context for SELinux-enabled systems.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"fscontext",
		"Sets the SELinux context for the filesystem itself. This is used to set the security context for the filesystem as a whole.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"defcontext",
		"Sets the default SELinux context for files that do not have a specific context set.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"rootcontext",
		"Sets the SELinux context for the root directory of the filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"aclinherit",
		"Controls how ACEs are inherited when files and directories are created.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("discard", "does not inherit any ACEs"),
			docvalues.CreateEnumStringWithDoc("noallow", "only inherits inheritable ACEs specifying deny permissions"),
			docvalues.CreateEnumStringWithDoc("restricted", "removes write_acl and write_owner permissions when inherited"),
			docvalues.CreateEnumStringWithDoc("passthrough", "inherits all inheritable ACEs without modifications"),
			docvalues.CreateEnumStringWithDoc("passthrough-x", "similar to passthrough, with execute permission nuances"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"aclmode",
		"Controls ACL modification during chmod operations.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("discard", "deletes all ACEs except those representing file mode"),
			docvalues.CreateEnumStringWithDoc("groupmask", "reduces permissions in ALLOW entries"),
			docvalues.CreateEnumStringWithDoc("passthrough", "no changes to ACL"),
			docvalues.CreateEnumStringWithDoc("restricted", "prevents chmod on files with non-trivial ACLs"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"acltype",
		"Controls ACL type and enablement.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("off", "ACLs disabled"),
			docvalues.CreateEnumStringWithDoc("nfsv4", "NFSv4-style ACLs"),
			docvalues.CreateEnumStringWithDoc("posix", "POSIX ACLs"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"checksum",
		"Controls data integrity verification algorithm.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "automatically selects algorithm"),
			docvalues.CreateEnumStringWithDoc("off", "disables integrity checking"),
			docvalues.CreateEnumStringWithDoc("fletcher2", "Fletcher2 checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("fletcher4", "Fletcher4 checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("sha256", "SHA256 checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("sha512", "SHA512 checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("skein", "Skein checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("edonr", "Edon-R checksum algorithm"),
			docvalues.CreateEnumStringWithDoc("blake3", "BLAKE3 checksum algorithm"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"compression",
		"Controls data compression algorithm.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "uses default compression"),
			docvalues.CreateEnumStringWithDoc("off", "no compression"),
			docvalues.CreateEnumStringWithDoc("gzip", "gzip compression"),
			docvalues.CreateEnumStringWithDoc("gzip-1", "gzip compression level 1"),
			docvalues.CreateEnumStringWithDoc("gzip-2", "gzip compression level 2"),
			docvalues.CreateEnumStringWithDoc("gzip-3", "gzip compression level 3"),
			docvalues.CreateEnumStringWithDoc("gzip-4", "gzip compression level 4"),
			docvalues.CreateEnumStringWithDoc("gzip-5", "gzip compression level 5"),
			docvalues.CreateEnumStringWithDoc("gzip-6", "gzip compression level 6"),
			docvalues.CreateEnumStringWithDoc("gzip-7", "gzip compression level 7"),
			docvalues.CreateEnumStringWithDoc("gzip-8", "gzip compression level 8"),
			docvalues.CreateEnumStringWithDoc("gzip-9", "gzip compression level 9"),
			docvalues.CreateEnumStringWithDoc("lz4", "LZ4 compression"),
			docvalues.CreateEnumStringWithDoc("lzjb", "LZJB compression"),
			docvalues.CreateEnumStringWithDoc("zle", "Zero Length Encoding"),
			docvalues.CreateEnumStringWithDoc("zstd", "Zstandard compression"),
			docvalues.CreateEnumStringWithDoc("zstd-1", "Zstandard compression level 1"),
			docvalues.CreateEnumStringWithDoc("zstd-2", "Zstandard compression level 2"),
			docvalues.CreateEnumStringWithDoc("zstd-3", "Zstandard compression level 3"),
			docvalues.CreateEnumStringWithDoc("zstd-4", "Zstandard compression level 4"),
			docvalues.CreateEnumStringWithDoc("zstd-5", "Zstandard compression level 5"),
			docvalues.CreateEnumStringWithDoc("zstd-6", "Zstandard compression level 6"),
			docvalues.CreateEnumStringWithDoc("zstd-7", "Zstandard compression level 7"),
			docvalues.CreateEnumStringWithDoc("zstd-8", "Zstandard compression level 8"),
			docvalues.CreateEnumStringWithDoc("zstd-9", "Zstandard compression level 9"),
			docvalues.CreateEnumStringWithDoc("zstd-10", "Zstandard compression level 10"),
			docvalues.CreateEnumStringWithDoc("zstd-11", "Zstandard compression level 11"),
			docvalues.CreateEnumStringWithDoc("zstd-12", "Zstandard compression level 12"),
			docvalues.CreateEnumStringWithDoc("zstd-13", "Zstandard compression level 13"),
			docvalues.CreateEnumStringWithDoc("zstd-14", "Zstandard compression level 14"),
			docvalues.CreateEnumStringWithDoc("zstd-15", "Zstandard compression level 15"),
			docvalues.CreateEnumStringWithDoc("zstd-16", "Zstandard compression level 16"),
			docvalues.CreateEnumStringWithDoc("zstd-17", "Zstandard compression level 17"),
			docvalues.CreateEnumStringWithDoc("zstd-18", "Zstandard compression level 18"),
			docvalues.CreateEnumStringWithDoc("zstd-19", "Zstandard compression level 19"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-1", "Zstandard fast compression level 1"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-2", "Zstandard fast compression level 2"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-3", "Zstandard fast compression level 3"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-4", "Zstandard fast compression level 4"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-5", "Zstandard fast compression level 5"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-6", "Zstandard fast compression level 6"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-7", "Zstandard fast compression level 7"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-8", "Zstandard fast compression level 8"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-9", "Zstandard fast compression level 9"),
			docvalues.CreateEnumStringWithDoc("zstd-fast-10", "Zstandard fast compression level 10"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"copies",
		"Controls the number of copies of data stored for this dataset. These copies are in addition to any redundancy provided by the pool.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("1", "single copy"),
			docvalues.CreateEnumStringWithDoc("2", "two copies"),
			docvalues.CreateEnumStringWithDoc("3", "three copies"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"dedup",
		"Controls whether deduplication is enabled for this dataset.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "enable deduplication"),
			docvalues.CreateEnumStringWithDoc("off", "disable deduplication"),
			docvalues.CreateEnumStringWithDoc("verify", "enable deduplication with verification"),
			docvalues.CreateEnumStringWithDoc("sha256", "enable deduplication using SHA256"),
			docvalues.CreateEnumStringWithDoc("sha256,verify", "enable deduplication using SHA256 with verification"),
			docvalues.CreateEnumStringWithDoc("sha512", "enable deduplication using SHA512"),
			docvalues.CreateEnumStringWithDoc("sha512,verify", "enable deduplication using SHA512 with verification"),
			docvalues.CreateEnumStringWithDoc("skein", "enable deduplication using Skein"),
			docvalues.CreateEnumStringWithDoc("skein,verify", "enable deduplication using Skein with verification"),
			docvalues.CreateEnumStringWithDoc("edonr", "enable deduplication using Edon-R"),
			docvalues.CreateEnumStringWithDoc("edonr,verify", "enable deduplication using Edon-R with verification"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"logbias",
		"Provides a hint to ZFS about handling of synchronous requests in this dataset.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("latency", "optimize for low latency"),
			docvalues.CreateEnumStringWithDoc("throughput", "optimize for high throughput"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"primarycache",
		"Controls what is cached in the primary cache (ARC).",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("all", "cache both metadata and data"),
			docvalues.CreateEnumStringWithDoc("none", "cache neither metadata nor data"),
			docvalues.CreateEnumStringWithDoc("metadata", "cache only metadata"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"quota",
		"Limits the amount of disk space a dataset and its descendents can consume.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"recordsize",
		"Specifies a suggested block size for files in the filesystem.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("512", "512 bytes"),
			docvalues.CreateEnumStringWithDoc("1K", "1 KB"),
			docvalues.CreateEnumStringWithDoc("2K", "2 KB"),
			docvalues.CreateEnumStringWithDoc("4K", "4 KB"),
			docvalues.CreateEnumStringWithDoc("8K", "8 KB"),
			docvalues.CreateEnumStringWithDoc("16K", "16 KB"),
			docvalues.CreateEnumStringWithDoc("32K", "32 KB"),
			docvalues.CreateEnumStringWithDoc("64K", "64 KB"),
			docvalues.CreateEnumStringWithDoc("128K", "128 KB"),
			docvalues.CreateEnumStringWithDoc("256K", "256 KB"),
			docvalues.CreateEnumStringWithDoc("512K", "512 KB"),
			docvalues.CreateEnumStringWithDoc("1M", "1 MB"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"refquota",
		"Limits the amount of space a dataset can consume. This hard limit does not include space used by descendents.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"refreservation",
		"The minimum amount of space guaranteed to a dataset, not including its descendents.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"reservation",
		"The minimum amount of space guaranteed to a dataset and its descendents.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"secondarycache",
		"Controls what is cached in the secondary cache (L2ARC).",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("all", "cache both metadata and data"),
			docvalues.CreateEnumStringWithDoc("none", "cache neither metadata nor data"),
			docvalues.CreateEnumStringWithDoc("metadata", "cache only metadata"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"snapdir",
		"Controls whether the .zfs directory is hidden or visible in the root of the filesystem.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("hidden", "hide .zfs directory"),
			docvalues.CreateEnumStringWithDoc("visible", "show .zfs directory"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"sync",
		"Controls the behavior of synchronous requests.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("standard", "honor synchronous requests"),
			docvalues.CreateEnumStringWithDoc("always", "every write is synchronous"),
			docvalues.CreateEnumStringWithDoc("disabled", "disable synchronous requests"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"version",
		"The on-disk version of this filesystem.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("1", "version 1"),
			docvalues.CreateEnumStringWithDoc("2", "version 2"),
			docvalues.CreateEnumStringWithDoc("3", "version 3"),
			docvalues.CreateEnumStringWithDoc("4", "version 4"),
			docvalues.CreateEnumStringWithDoc("5", "version 5"),
			docvalues.CreateEnumStringWithDoc("current", "current version"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"volsize",
		"For volumes, specifies the logical size of the volume.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"xattr",
		"Controls whether extended attributes are enabled for this filesystem.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "enable extended attributes using system attribute directory"),
			docvalues.CreateEnumStringWithDoc("off", "disable extended attributes"),
			docvalues.CreateEnumStringWithDoc("dir", "enable extended attributes using directory-based storage"),
			docvalues.CreateEnumStringWithDoc("sa", "enable extended attributes using system attribute storage"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"zoned",
		"Controls whether the dataset is managed from a non-global zone.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "dataset is managed from a non-global zone"),
			docvalues.CreateEnumStringWithDoc("off", "dataset is managed from the global zone"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"redundant_metadata",
		"Controls what types of metadata are stored redundantly.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("all", "store all metadata redundantly"),
			docvalues.CreateEnumStringWithDoc("most", "store most metadata redundantly"),
			docvalues.CreateEnumStringWithDoc("some", "store some metadata redundantly"),
			docvalues.CreateEnumStringWithDoc("none", "store no metadata redundantly"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"special_small_blocks",
		"Threshold block size for including small file or zvol blocks into special allocation class.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"filesystem_limit",
		"Limits number of filesystems and volumes under this point in dataset tree.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"snapshot_limit",
		"Limits number of snapshots that can be created on a dataset.",
	): numericProp,
	docvalues.CreateEnumStringWithDoc(
		"canmount",
		"Controls whether file system can be mounted.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "filesystem can be mounted"),
			docvalues.CreateEnumStringWithDoc("off", "filesystem cannot be mounted"),
			docvalues.CreateEnumStringWithDoc("noauto", "filesystem can be mounted but not automatically"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"mountpoint",
		"Controls mount point used for file system.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"sharenfs",
		"Controls NFS sharing and options.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"sharesmb",
		"Controls SMB sharing.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"dnodesize",
		"Specifies size of dnodes in file system.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("legacy", "use legacy dnode size"),
			docvalues.CreateEnumStringWithDoc("auto", "automatically select dnode size"),
			docvalues.CreateEnumStringWithDoc("1k", "1 KB dnode size"),
			docvalues.CreateEnumStringWithDoc("2k", "2 KB dnode size"),
			docvalues.CreateEnumStringWithDoc("4k", "4 KB dnode size"),
			docvalues.CreateEnumStringWithDoc("8k", "8 KB dnode size"),
			docvalues.CreateEnumStringWithDoc("16k", "16 KB dnode size"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"encryption",
		"Controls encryption cipher suite.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("off", "no encryption"),
			docvalues.CreateEnumStringWithDoc("on", "use default encryption"),
			docvalues.CreateEnumStringWithDoc("aes-128-ccm", "AES-128-CCM encryption"),
			docvalues.CreateEnumStringWithDoc("aes-192-ccm", "AES-192-CCM encryption"),
			docvalues.CreateEnumStringWithDoc("aes-256-ccm", "AES-256-CCM encryption"),
			docvalues.CreateEnumStringWithDoc("aes-128-gcm", "AES-128-GCM encryption"),
			docvalues.CreateEnumStringWithDoc("aes-192-gcm", "AES-192-GCM encryption"),
			docvalues.CreateEnumStringWithDoc("aes-256-gcm", "AES-256-GCM encryption"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"keylocation",
		"Controls where encryption key is loaded from.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"keyformat",
		"Controls encryption key format.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("raw", "raw key format"),
			docvalues.CreateEnumStringWithDoc("hex", "hexadecimal key format"),
			docvalues.CreateEnumStringWithDoc("passphrase", "passphrase key format"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"pbkdf2iters",
		"Controls the number of PBKDF2 iterations used for generating encryption keys from passphrases.",
	): docvalues.NumberRangeValue(100000, 1000000000),
	docvalues.CreateEnumStringWithDoc(
		"devices",
		"Controls whether device nodes can be opened on this file system.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "allow device nodes"),
			docvalues.CreateEnumStringWithDoc("off", "disallow device nodes"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"setuid",
		"Controls whether the setuid bit is respected for the file system.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "respect setuid bits"),
			docvalues.CreateEnumStringWithDoc("off", "ignore setuid bits"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"readonly",
		"Controls whether this dataset can be modified.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "read-only access"),
			docvalues.CreateEnumStringWithDoc("off", "read-write access"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"atime",
		"Controls whether file access times are updated when they are read.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "update access times"),
			docvalues.CreateEnumStringWithDoc("off", "do not update access times"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"relatime",
		"Controls the manner in which the access time is updated.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "update access time relative to modify time"),
			docvalues.CreateEnumStringWithDoc("off", "disable relative access time updates"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"nbmand",
		"Controls whether the file system should be mounted with non-blocking mandatory locks.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "enable non-blocking mandatory locks"),
			docvalues.CreateEnumStringWithDoc("off", "disable non-blocking mandatory locks"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"vscan",
		"Controls whether regular files should be scanned for viruses when a file is opened and closed. In addition to enabling this property, the virus scan service must also be enabled for virus scanning to occur. The default value is off. This property is not used by OpenZFS.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "enable virus scanning"),
			docvalues.CreateEnumStringWithDoc("off", "disable virus scanning"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"jailed",
		"Controls whether the dataset is managed from a jail. See zfs-jail(8) for more information. Jails are a FreeBSD feature and this property is not available on other platforms.",
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("on", "dataset is managed from a jail"),
			docvalues.CreateEnumStringWithDoc("off", "dataset is not managed from a jail"),
		},
	},
}

var ZfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"dev",
		"Controls whether device nodes can be opened on this file system. This is the default behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodev",
		"Controls whether device nodes can be opened on this file system. When set, device nodes cannot be opened.",
	),
	docvalues.CreateEnumStringWithDoc(
		"exec",
		"Controls whether processes can be executed from within this file system. This is the default behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noexec",
		"Controls whether processes can be executed from within this file system. When set, executables cannot be run from this filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"ro",
		"Controls whether this dataset can be modified. When set, the filesystem is mounted read-only.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rw",
		"Controls whether this dataset can be modified. This is the default behavior allowing read-write access.",
	),
	docvalues.CreateEnumStringWithDoc(
		"suid",
		"Controls whether the setuid bit is respected for the file system. This is the default behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nosuid",
		"Controls whether the setuid bit is respected for the file system. When set, setuid and setgid bits are ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"devices",
		"Controls whether device nodes can be opened on this file system. This is the default behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodevices",
		"Controls whether device nodes can be opened on this file system. When set, device nodes cannot be opened.",
	),
	docvalues.CreateEnumStringWithDoc(
		"setuid",
		"Controls whether the setuid bit is respected for the file system. This is the default behavior.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nosetuid",
		"Controls whether the setuid bit is respected for the file system. When set, setuid and setgid bits are ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"readonly",
		"Controls whether this dataset can be modified. When set, the filesystem is mounted read-only.",
	),
	docvalues.CreateEnumStringWithDoc(
		"readwrite",
		"Controls whether this dataset can be modified. This is the default behavior allowing read-write access.",
	),
	docvalues.CreateEnumStringWithDoc(
		"overlay",
		"Allow mounting on busy or populated directory.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nooverlay",
		"Disallow mounting on busy or populated directory.",
	),
}
