package fields

import docvalues "config-lsp/doc-values"

var FsckField = docvalues.EnumValue{
	EnforceValues: false,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumStringWithDoc(
			"0",
			"Defaults to zero (don’t check the filesystem) if not present.",
		),
		docvalues.CreateEnumStringWithDoc(
			"1",
			"The root filesystem should be specified with a fs_passno of 1.",
		),
		docvalues.CreateEnumStringWithDoc(
			"2",
			"Other filesystems [than the root filesystem] should have a fs_passno of 2.",
		),
	},
}

var FsckFieldWhenDisabledFilesystems = docvalues.EnumValue{
	EnforceValues: false,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumStringWithDoc(
			"0",
			"Defaults to zero (don’t check the filesystem) if not present.",
		),
		docvalues.CreateEnumStringWithDoc(
			"2",
			"Other filesystems [than the root filesystem] should have a fs_passno of 2.",
		),
	},
}

var FsckOneDisabledFilesystems = map[string]struct{}{
	"btrfs": {},
	"xfs":   {},
}
