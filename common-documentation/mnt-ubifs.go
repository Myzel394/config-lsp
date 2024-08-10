package commondocumentation

import docvalues "config-lsp/doc-values"

var UbifsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"compr",
		"Select the default compressor which is used when new files are written. It is still possible to read compressed files if mounted with the none option.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("none"),
			docvalues.CreateEnumString("lzo"),
			docvalues.CreateEnumString("zlib"),
		},
	},
}

var UbifsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"bulk_read",
		"Enable bulk-read. VFS read-ahead is disabled because it slows down the filesystem. Bulk-Read is an internal optimization. Some flashes may read faster if the data are read at one go, rather than at several read requests. For example, OneNAND can do \"read-while-load\" if it reads more than one NAND page.",
	),
	docvalues.CreateEnumStringWithDoc(
		"no_bulk_read",
		"Do not bulk-read. This is the default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"chk_data_crc",
		"Check data CRC-32 checksums. This is the default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"no_chk_data_crc",
		"Do not check data CRC-32 checksums. With this option, the filesystem does not check CRC-32 checksum for data, but it does check it for the internal indexing information. This option only affects reading, not writing. CRC-32 is always calculated when writing the data.",
	),
}
