package commondocumentation

import (
	docvalues "config-lsp/doc-values"
)

var FatDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"blocksize",
		"Set blocksize (default 512). This option is obsolete.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("512"),
			docvalues.CreateEnumString("1024"),
			docvalues.CreateEnumString("2048"),
		},
	},
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
		"umask",
		"Set the umask (the bitmask of the permissions that are not present). The default is the umask of the current process. The value is given in octal.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"dmask",
		"Set the umask applied to directories only. The default is the umask of the current process. The value is given in octal.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"fmask",
		"Set the umask applied to regular files only. The default is the umask of the current process. The value is given in octal.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"allow_utime",
		"This option controls the permission check of mtime/atime.",
	): docvalues.UmaskValue{},
	docvalues.CreateEnumStringWithDoc(
		"check",
		"Three different levels of pickiness can be chosen: relaxed, normal, strict.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("relaxed"),
			docvalues.CreateEnumString("normal"),
			docvalues.CreateEnumString("strict"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"codepage",
		"Sets the codepage for converting to shortname characters on FAT and VFAT filesystems. By default, codepage 437 is used.",
	): docvalues.NumberValue{Min: &zero},
	docvalues.CreateEnumStringWithDoc(
		// TODO: Show warning in analyzevalue when used
		"conv",
		"This option is obsolete and may fail or be ignored.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"cvf_format",
		"Forces the driver to use the CVF (Compressed Volume File) module cvf_module instead of auto-detection. If the kernel supports kmod, the cvf_format=xxx option also controls on-demand CVF module loading. This option is obsolete.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"cvf_option",
		"Option passed to the CVF module. This option is obsolete.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"errors",
		"Specify FAT behavior on critical errors: panic, continue without doing anything, or remount the partition in read-only mode (default behavior).",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"panic",
				"Causes the kernel to panic on errors.",
			),
			docvalues.CreateEnumStringWithDoc(
				"continue",
				"Continues without doing anything.",
			),
			docvalues.CreateEnumStringWithDoc(
				"remount",
				"Remounts the partition in read-only mode.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"fat",
		"Specify a 12, 16 or 32 bit fat. This overrides the automatic FAT type detection routine. Use with caution!",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("12"),
			docvalues.CreateEnumString("16"),
			docvalues.CreateEnumString("32"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Character set to use for converting between 8 bit characters and 16 bit Unicode characters. The default is iso8859-1. Long filenames are stored on disk in Unicode format.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
	docvalues.CreateEnumStringWithDoc(
		"nfs",
		`Enable this only if you want to export the FAT filesystem over NFS.
	To maintain backward compatibility, -o nfs is also accepted, defaulting to stale_rw`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"stale_rw",
				"This option maintains an index (cache) of directory inodes which is used by the nfs-related code to improve look-ups. Full file operations (read/write) over NFS are supported but with cache eviction at NFS server, this could result in spurious ESTALE errors.",
			),
			docvalues.CreateEnumStringWithDoc(
				"nostale_ro",
				"This option bases the inode number and file handle on the on-disk location of a file in the FAT directory entry. This ensures that ESTALE will not be returned after a file is evicted from the inode cache. However, it means that operations such as rename, create and unlink could cause file handles that previously pointed at one file to point at a different file, potentially causing data corruption. For this reason, this option also mounts the filesystem readonly.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"tz",
		"This option disables the conversion of timestamps between local time (as used by Windows on FAT) and UTC (which Linux uses internally). This is particularly useful when mounting devices (like digital cameras) that are set to UTC in order to avoid the pitfalls of local time.",
		// TODO: Add enum for timezones
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"time_offset",
		"Set offset for conversion of timestamps from local time used by FAT to UTC. I.e., minutes will be subtracted from each timestamp to convert it to UTC used internally by Linux. This is useful when the time zone set in the kernel via settimeofday(2) is not the time zone used by the filesystem. Note that this option still does not provide correct time stamps in all cases in presence of DST - time stamps in a different DST setting will be off by one hour.",
	): docvalues.NumberValue{},
	docvalues.CreateEnumStringWithDoc(
		"dotsOK",
		"Various misguided attempts to force Unix or DOS conventions onto a FAT filesystem.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("yes"),
			docvalues.CreateEnumString("no"),
		},
	},
}

var FatDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"debug",
		"Turn on the debug flag. A version string and a list of filesystem parameters will be printed (these data are also printed if the parameters appear to be inconsistent).",
	),
	docvalues.CreateEnumStringWithDoc(
		"discard",
		"If set, causes discard/TRIM commands to be issued to the block device when blocks are freed. This is useful for SSD devices and sparse/thinly-provisioned LUNs.",
	),
	docvalues.CreateEnumStringWithDoc(
		"dos1xfloppy",
		"If set, use a fallback default BIOS Parameter Block configuration, determined by backing device size. These static parameters match defaults assumed by DOS 1.x for 160 kiB, 180 kiB, 320 kiB, and 360 kiB floppies and floppy images.",
	),
	docvalues.CreateEnumStringWithDoc(
		"quiet",
		"Turn on the quiet flag. Attempts to chown or chmod files do not return errors, although they fail. Use with caution!",
	),
	docvalues.CreateEnumStringWithDoc(
		"rodir",
		"FAT has the ATTR_RO (read-only) attribute. On Windows, the ATTR_RO of the directory will just be ignored, and is used only by applications as a flag (e.g. it’s set for the customized folder). If you want to use ATTR_RO as read-only flag even for the directory, set this option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"showexec",
		"If set, the execute permission bits of the file will be allowed only if the extension part of the name is .EXE, .COM, or .BAT. Not set by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"sys_immutable",
		"If set, ATTR_SYS attribute on FAT is handled as IMMUTABLE flag on Linux. Not set by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"flush",
		"If set, the filesystem will try to flush to disk more early than normal. Not set by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"usefree",
		"Use the \"free clusters\" value stored on FSINFO. It’ll be used to determine number of free clusters without scanning disk. But it’s not used by default, because recent Windows don’t update it correctly in some case. If you are sure the \"free clusters\" on FSINFO is correct, by this option you can avoid scanning disk.",
	),
	docvalues.CreateEnumStringWithDoc(
		"dots",
		"Various misguided attempts to force Unix or DOS conventions onto a FAT filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodots",
		"Various misguided attempts to force Unix or DOS conventions onto a FAT filesystem.",
	),
}
