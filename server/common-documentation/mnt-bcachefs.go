package commondocumentation

import docvalues "config-lsp/doc-values"

var checksumType = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumString("none"),
		docvalues.CreateEnumString("crc32c"),
		docvalues.CreateEnumString("crc64"),
	},
}

var compressionType = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumStringWithDoc("none", "(default)"),
		docvalues.CreateEnumString("lz4"),
		docvalues.CreateEnumString("gzip"),
		docvalues.CreateEnumString("zstd"),
	},
}

// No idea if those enums are correct,
// the documentation does not provide any information
var booleanEnumValue = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumString("yes"),
		docvalues.CreateEnumString("no"),
	},
}

var BcacheFSDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"errors",
		"Action to take on filesystem error. The errors option is used for inconsistencies that indicate some sort of a bug",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("continue", "Log the error but continue normal operation"),
			docvalues.CreateEnumStringWithDoc("ro", "Emergency read only, immediately halting any changes to the filesystem on disk"),
			docvalues.CreateEnumStringWithDoc("panic", "Immediately halt the entire machine, printing a backtrace on the system console"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"metadata_replicas",
		"Number of replicas for metadata (journal and btree)",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"data_replicas",
		"Number of replicas for user data",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"metadata_checksum",
		"Checksum type for metadata writes",
	): checksumType,
	docvalues.CreateEnumStringWithDoc(
		"data_checksum",
		"Checksum type for data writes",
	): checksumType,
	docvalues.CreateEnumStringWithDoc(
		"compression",
		"Compression type",
	): compressionType,
	docvalues.CreateEnumStringWithDoc(
		"background_compression",
		"Background compression type",
	): compressionType,
	docvalues.CreateEnumStringWithDoc(
		"str_hash",
		"Hash function for string hash tables (directories and xattrs)",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("crc32c"),
			docvalues.CreateEnumString("crc64"),
			docvalues.CreateEnumString("siphash"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"metadata_target",
		"Preferred target for metadata writes",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"foreground_target",
		"Preferred target for foreground writes",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"background_target",
		"Target for data to be moved to in the background",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"promote_target",
		"Target for data to be copied to on read",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"erasure_code",
		"Enable erasure coding",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"inodes_32bit",
		"Restrict new inode numbers to 32 bits",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"shard_inode_numbers",
		"Use CPU id for high bits of new inode numbers.",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"wide_macs",
		"Store full 128 bit cryptographic MACs (default 80)",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"inline_data",
		"Enable inline data extents (default on)",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"journal_flush_delay",
		"Delay in milliseconds before automatic journal commit (default 1000)",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"journal_flush_disabled",
		"Disables journal flush on sync/fsync. `journal_flush_delay` remains in effect, thus with the default setting not more than 1 second of work will be lost",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"journal_reclaim",
		"Reclaim journal space after a certain amount of time",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"journal_reclaim_delay",
		"Delay in milliseconds before automatic journal reclaim",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"acl",
		"Enable POSIX ACLs",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"usrquota",
		"Enable user quotas",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"grpquota",
		"Enable group quotas",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"prjquota",
		"Enable project quotas",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"degraded",
		"Allow mounting with data degraded",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"very_degraded",
		"Allow mounting with data missing",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"verbose",
		"Extra debugging info during mount/recovery",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"fsck",
		"Run fsck during mount",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"fix_errors",
		"Fix errors without asking during fsck",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"ratelimit_errors",
		"Ratelimit error messages during fsck",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"read_only",
		"Mount in read only mode",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"nochanges",
		"Issue no writes, even for journal replay",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"norecovery",
		"Don’t replay the journal (not recommended)",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"noexcl",
		"Don’t open devices in exclusive mode",
	): booleanEnumValue,
	docvalues.CreateEnumStringWithDoc(
		"version_upgrade",
		"Upgrade on disk format to latest version",
	): booleanEnumValue,
}

var BcacheFSDocumentationEnums = []docvalues.EnumString{}
