package commondocumentation

import docvalues "config-lsp/doc-values"

var FatDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"blocksize",
		"Set blocksize (default 512). This option is obsolete.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			// TODO: Check for other values too
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
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"dmask",
		"Set the umask applied to directories only. The default is the umask of the current process. The value is given in octal.",
		// TODO: Add mask
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"fmask",
		"Set the umask applied to regular files only. The default is the umask of the current process. The value is given in octal.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"allow_utime",
		"This option controls the permission check of mtime/atime.",
	): docvalues.StringValue{},
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
		// TODO: Check if NumberValue fits here 
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		"Character set to use for converting between 8 bit characters and 16 bit Unicode characters. The default is iso8859-1. Long filenames are stored on disk in Unicode format.",
	// TODO: Use enum for charsets
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"time_offset",
		"Set offset for conversion of timestamps from local time used by FAT to UTC. I.e., minutes will be subtracted from each timestamp to convert it to UTC used internally by Linux. This is useful when the time zone set in the kernel via settimeofday(2) is not the time zone used by the filesystem. Note that this option still does not provide correct time stamps in all cases in presence of DST - time stamps in a different DST setting will be off by one hour.",
		// TODO: Probably NumberValeu
	): docvalues.NumberValue{},
}

var FatDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		// TODO: Show warning in analyzevalue when used
		"conv",
		"This option is obsolete and may fail or be ignored.",
	),
	docvalues.CreateEnumStringWithDoc(
		"cvf_format",
		"Forces the driver to use the CVF (Compressed Volume File) module cvf_module instead of auto-detection. If the kernel supports kmod, the cvf_format=xxx option also controls on-demand CVF module loading. This option is obsolete.",
	),
	docvalues.CreateEnumStringWithDoc(
		"cvf_option",
		"Option passed to the CVF module. This option is obsolete.",
	),
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
	// TODO: Should be in map
	docvalues.CreateEnumStringWithDoc(
		"errors",
		"Specify FAT behavior on critical errors: panic, continue without doing anything, or remount the partition in read-only mode (default behavior).",
	),
	docvalues.CreateEnumStringWithDoc(
		"fat",
		"Specify a 12, 16 or 32 bit fat. This overrides the automatic FAT type detection routine. Use with caution!",
	),
	docvalues.CreateEnumStringWithDoc(
		"nfs",
		"Enable this only if you want to export the FAT filesystem over NFS.",
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
	docvalues.CreateEnumStringWithDoc(
		"dotsOK",
		"Various misguided attempts to force Unix or DOS conventions onto a FAT filesystem.",
	),
}
