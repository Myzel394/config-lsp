// This file contains documentation for ext2 filesystem options
package commondocumentation

import docvalues "config-lsp/doc-values"

var Ext2DocumentationAssignable = docvalues.KeyEnumAssignmentValue{
	Separator: ",",
	ValueIsOptional: false,
	Values: map[docvalues.EnumString]docvalues.Value{
		docvalues.CreateEnumStringWithDoc(
			"none",
			"No checking is done at mount time. This is the default. This is fast.  It is wise to invoke e2fsck(8) every now and then, e.g. at boot time. The non-default behavior is unsupported (check=normal and check=strict options have been removed). Note that these mount options don't have to be supported if ext4 kernel driver is used for ext2 and ext3 file systems.",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc(
					"none",
					"No checking is done at mount time",
				),
			},
		},
		docvalues.CreateEnumStringWithDoc(
			"errors",
			"Define the behavior when an error is encountered.  (Either ignore errors and just mark the file system erroneous and continue, or remount the file system read-only, or panic and halt the system.)  The default is set in the file system superblock, and can be changed using tune2fs(8).",
		): docvalues.EnumValue{
			EnforceValues: true,
			Values: []docvalues.EnumString{
				docvalues.CreateEnumStringWithDoc(
					"continue",
					"Ignore errors and just mark the file system erroneous and continue",
				),
				docvalues.CreateEnumStringWithDoc(
					"remount-ro",
					"Remount the file system read-only",
				),
				docvalues.CreateEnumStringWithDoc(
					"panic",
					"Panic and halt the system",
				),
			},
		},
		docvalues.CreateEnumStringWithDoc(
			"resgid",
			"The ext2 file system reserves a certain percentage of the available space (by default 5%, see mke2fs(8) and tune2fs(8)).  These options determine who can use the reserved blocks.  (Roughly: whoever has the specified uid, or belongs to the specified group.)",
		): docvalues.PositiveNumberValue(),
		docvalues.CreateEnumStringWithDoc(
			"resuid",
			"The ext2 file system reserves a certain percentage of the available space (by default 5%, see mke2fs(8) and tune2fs(8)).  These options determine who can use the reserved blocks.  (Roughly: whoever has the specified uid, or belongs to the specified group.)",
		): docvalues.PositiveNumberValue(),
		docvalues.CreateEnumStringWithDoc(
			"sb",
			"Instead of using the normal superblock, use an alternative superblock specified by n.  This option is normally used when the primary superblock has been corrupted.  The location of backup superblocks is dependent on the file system's blocksize, the number of blocks per group, and features such as sparse_super.",
		): docvalues.PositiveNumberValue(),

	},
}

var Ext2DocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"acl",
		"Enable POSIX Access Control Lists (ACLs). See the acl(5) manual page.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noacl",
		"Disable POSIX Access Control Lists (ACLs). See the acl(5) manual page.",
	),
	docvalues.CreateEnumStringWithDoc(
		"bsddf",
		`Set the behavior for the statfs system call. The minixdf behavior is to return in the f_blocks field the total number of blocks of the file system, while the bsddf behavior (which is the default) is to subtract the overhead blocks used by the ext2 file system and not available for file storage. Thus

	% mount /k -o minixdf; df /k; umount /k
	File System  1024-blocks   Used  Available  Capacity  Mounted on
	/dev/sda6      2630655    86954   2412169      3%     /k

	% mount /k -o bsddf; df /k; umount /k
	File System  1024-blocks  Used  Available  Capacity  Mounted on
	/dev/sda6      2543714      13   2412169      0%     /k

	(Note that this example shows that one can add command
	line options to the options given in /etc/fstab.)`,
	),
	docvalues.CreateEnumStringWithDoc(
		"minixdf",
		`Set the behavior for the statfs system call. The minixdf behavior is to return in the f_blocks field the total number of blocks of the file system, while the bsddf behavior (which is the default) is to subtract the overhead blocks used by the ext2 file system and not available for file storage. Thus

	% mount /k -o minixdf; df /k; umount /k
	File System  1024-blocks   Used  Available  Capacity  Mounted on
	/dev/sda6      2630655    86954   2412169      3%     /k

	% mount /k -o bsddf; df /k; umount /k
	File System  1024-blocks  Used  Available  Capacity  Mounted on
	/dev/sda6      2543714      13   2412169      0%     /k

	(Note that this example shows that one can add command
	line options to the options given in /etc/fstab.)`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nocheck",
		"No checking is done at mount time. This is the default. This is fast.  It is wise to invoke e2fsck(8) every now and then, e.g. at boot time. The non-default behavior is unsupported (check=normal and check=strict options have been removed). Note that these mount options don't have to be supported if ext4 kernel driver is used for ext2 and ext3 file systems.",
	),
	docvalues.CreateEnumStringWithDoc(
		"debug",
		"Print debugging info upon each (re)mount.",
	),

	docvalues.CreateEnumStringWithDoc(
		"grpid",
		"Newly created files will take on the group of the directory which they are created in.",
	),
	docvalues.CreateEnumStringWithDoc(
		"bsdgroups",
		"Take the fsgid of the current process, unless the directory has the setgid bit set, in which case it takes the gid from the parent directory, and also gets the setgid bit set if it is a directory itself",
	),
	docvalues.CreateEnumString(
		"nogrpid",
	),
	docvalues.CreateEnumString(
		"sysvgroups",
	),

	docvalues.CreateEnumStringWithDoc(
		"grpquota",
		"grpquota enables group quotas support. You need the quota utilities to actually enable and manage the quota system.",
	),
	docvalues.CreateEnumString(
		"noquota",
	),
	docvalues.CreateEnumStringWithDoc(
		"quota",
		"The usrquota (same as quota) mount option enables user quota support on the file system.",
	),
	docvalues.CreateEnumStringWithDoc(
		"usrquota",
		"The usrquota (same as quota) mount option enables user quota support on the file system.",
	),

	docvalues.CreateEnumStringWithDoc(
		"nouid32",
		"Disables 32-bit UIDs and GIDs.  This is for interoperability with older kernels which only store and expect 16-bit values.",
	),

	docvalues.CreateEnumStringWithDoc(
		"oldalloc",
		"Use old allocator allocator for new inodes",
	),
	docvalues.CreateEnumStringWithDoc(
		"orlov",
		"Use Orlov block allocator for new inodes (default)",
	),
	docvalues.CreateEnumStringWithDoc(
		"user_xattr",
		"Support \"user.\" extended attributes",
	),
	docvalues.CreateEnumStringWithDoc(
		"nouser_xattr",
		"Do not support \"user.\" extended attributes",
	),
}
