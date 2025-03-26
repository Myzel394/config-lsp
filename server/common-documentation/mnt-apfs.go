package commondocumentation

import (
	docvalues "config-lsp/doc-values"
)

var APFSDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"user",
		"Set the owner of the files in the file system to user. The default owner is the owner of the directory on which the file system is being mounted. The user may be a user-name, or a numeric value.",
	): docvalues.UIDValue{},
	docvalues.CreateEnumStringWithDoc(
		"group",
		"Set the group of the files in the file system to group. The default group is the group of the directory on which the file system is being mounted. The group may be a group-name, or a numeric value.",
	): docvalues.GIDValue{},
	docvalues.CreateEnumStringWithDoc(
		"snapshot",
		"The name of the snapshot to mount. In this usage pathname is the mounted root directory of the base volume containing the snapshot.",
	): docvalues.StringValue{},
}

var APFSDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"async",
		"All I/O to the file system should be done asynchronously. This can be somewhat dangerous with respect to losing data when faced with system crashes and power outages. This is also the default. It can be avoided with the noasync option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noauto",
		"This filesystem should be skipped when mount is run with the -a flag.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodev",
		"Do not interpret character or block special devices on the file system. This option is useful for a server that has file systems containing special devices for architectures other than its own.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noexec",
		"Do not allow execution of any binaries on the mounted file system. This option is useful for a server that has file systems containing binaries for architectures other than its own.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noowners",
		"Ignore the ownership field for the entire volume. This causes all objects to appear as owned by user ID 99 and group ID 99. User ID 99 is interpreted as the current effective user ID, while group ID 99 is used directly and translates to ``unknown''.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nosuid",
		"Do not allow set-user-identifier or set-group-identifier bits to take effect.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rdonly",
		"The same as -r; mount the file system read-only (even the super-user may not write it).",
	),
	docvalues.CreateEnumStringWithDoc(
		"update",
		"The same as -u; indicate that the status of an already mounted file system should be changed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"union",
		"Causes the namespace to appear as the union of directories of the mounted filesystem with corresponding directories in the underlying filesystem. Lookups will be done in the mounted filesystem first. If those operations fail due to a non-existent file the underlying directory is then accessed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noatime",
		"Do not update the file access time when reading from a file. This option is useful on file systems where there are large numbers of files and performance is more critical than updating the file access time (which is rarely ever important).",
	),
	docvalues.CreateEnumStringWithDoc(
		"strictatime",
		"Always update the file access time when reading from a file. Without this option the filesystem may default to a less strict update mode, where some access time updates are skipped for performance reasons. This option could be ignored if it is not supported by the filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nobrowse",
		"This option indicates that the mount point should not be visible via the GUI (i.e., appear on the Desktop as a separate volume).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nofollow",
		"This option indicates that in the course of the mount system call, the kernel should not follow any symlinks that may be present in the provided mount-on directory. This is the same as the -k option.",
	),
}

