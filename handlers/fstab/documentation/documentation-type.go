package fstabdocumentation

import docvalues "config-lsp/doc-values"

var FileSystemTypeField = docvalues.ArrayValue{
	Separator:           ",",
	DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
	SubValue: docvalues.EnumValue{
		EnforceValues: false,
		Values: []docvalues.EnumString{
			{
				InsertText:      "none",
				DescriptionText: "none",
				Documentation:   "An entry _none_ is useful for bind or move mounts.",
			},
			{
				InsertText:      "swap",
				DescriptionText: "swap",
				Documentation:   "An entry _swap_ denotes a file or partition to be used for swapping, cf. swapon(8)",
			},
			{
				InsertText:      "ext2",
				DescriptionText: "ext2",
				Documentation:   "Mount as ext2 filesystem",
			},
			{
				InsertText:      "ext3",
				DescriptionText: "ext3",
				Documentation:   "Mount as ext2 filesystem",
			},
			{
				InsertText:      "ext4",
				DescriptionText: "ext4",
				Documentation:   "Mount as ext4 filesystem",
			},
			{
				InsertText:      "xfs",
				DescriptionText: "xfs",
				Documentation:   "Mount as xfs filesystem",
			},
			{
				InsertText:      "btrfs",
				DescriptionText: "btrfs",
				Documentation:   "Mount as btrfs filesystem",
			},
			{
				InsertText:      "f2fs",
				DescriptionText: "f2fs",
				Documentation:   "Mount as f2fs filesystem",
			},
			{
				InsertText:      "vfat",
				DescriptionText: "vfat",
				Documentation:   "Mount as vfat filesystem",
			},
			{
				InsertText:      "ntfs",
				DescriptionText: "ntfs",
				Documentation:   "Mount as ntfs filesystem",
			},
			{
				InsertText:      "hfsplus",
				DescriptionText: "hfsplus",
				Documentation:   "Mount as hfsplus filesystem",
			},
			{
				InsertText:      "tmpfs",
				DescriptionText: "tmpfs",
				Documentation:   "Mount as tmpfs filesystem",
			},
			{
				InsertText:      "sysfs",
				DescriptionText: "sysfs",
				Documentation:   "Mount as sysfs filesystem",
			},
			{
				InsertText:      "proc",
				DescriptionText: "proc",
				Documentation:   "Mount as proc filesystem",
			},
			{
				InsertText:      "iso9660",
				DescriptionText: "iso9660",
				Documentation:   "Mount as iso9660 filesystem",
			},
			{
				InsertText:      "udf",
				DescriptionText: "udf",
				Documentation:   "Mount as udf filesystem",
			},
			{
				InsertText:      "squashfs",
				DescriptionText: "squashfs",
				Documentation:   "Mount as squashfs filesystem",
			},
			{
				InsertText:      "nfs",
				DescriptionText: "nfs",
				Documentation:   "Mount as nfs filesystem",
			},
			{
				InsertText:      "cifs",
				DescriptionText: "cifs",
				Documentation:   "Mount as cifs filesystem",
			},
		},
	},
}
