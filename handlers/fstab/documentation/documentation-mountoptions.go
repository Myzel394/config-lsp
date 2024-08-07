package fstabdocumentation

import (
	commondocumentation "config-lsp/common-documentation/filesystems/mountoptions"
	docvalues "config-lsp/doc-values"
	"strings"
)

var mountOptionsExtractor = func(value string) string {
	separatorIndex := strings.Index(value, "=")

	if separatorIndex == -1 {
		return value
	}

	return value[:separatorIndex]
}

// From https://www.man7.org/linux/man-pages/man8/mount.8.html
var defaultOptions = []docvalues.EnumString{
	// Default options
	docvalues.CreateEnumStringWithDoc(
		"async",
		"All I/O to the filesystem should be done asynchronously. (See also the sync option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"atime",
		"Do not use the noatime feature, so the inode access time is controlled by kernel defaults. See also the descriptions of the relatime and strictatime mount options.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noatime",
		"Do not update inode access times on this filesystem (e.g. for faster access on the news spool to speed up news servers). This works for all inode types (directories too), so it implies nodiratime.",
	),
	docvalues.CreateEnumStringWithDoc(
		"auto",
		"Can be mounted with the -a option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noauto",
		"Can only be mounted explicitly (i.e., the -a option will not cause the filesystem to be mounted).",
	),
	docvalues.CreateEnumStringWithDoc(
		"context",
		"The context= option is useful when mounting filesystems that do not support extended attributes, such as a floppy or hard disk formatted with VFAT, or systems that are not normally running under SELinux, such as an ext3 or ext4 formatted disk from a non-SELinux workstation. You can also use context= on filesystems you do not trust, such as a floppy. It also helps in compatibility with xattr-supporting filesystems on earlier 2.4.<x> kernel versions. Even where xattrs are supported, you can save time not having to label every file by assigning the entire disk one security context.",
	),
	docvalues.CreateEnumStringWithDoc(
		"fscontext",
		"The fscontext= option works for all filesystems, regardless of their xattr support. The fscontext option sets the overarching filesystem label to a specific security context. This filesystem label is separate from the individual labels on the files. It represents the entire filesystem for certain kinds of permission checks, such as during mount or file creation. Individual file labels are still obtained from the xattrs on the files themselves. The context option actually sets the aggregate context that fscontext provides, in addition to supplying the same label for individual files.",
	),
	docvalues.CreateEnumStringWithDoc(
		"defcontext",
		"You can set the default security context for unlabeled files using defcontext= option. This overrides the value set for unlabeled files in the policy and requires a filesystem that supports xattr labeling.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rootcontext",
		"The rootcontext= option allows you to explicitly label the root inode of a FS being mounted before that FS or inode becomes visible to userspace. This was found to be useful for things like stateless Linux. The special value @target can be used to assign the current context of the target mountpoint location.",
	),
	docvalues.CreateEnumStringWithDoc(
		"defaults",
		"Use the default options: rw, suid, dev, exec, auto, nouser, and async. Note that the real set of all default mount options depends on the kernel and filesystem type. See the beginning of this section for more details.",
	),
	docvalues.CreateEnumStringWithDoc(
		"dev",
		"Interpret character or block special devices on the filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodev",
		"Do not interpret character or block special devices on the filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"diratime",
		"Update directory inode access times on this filesystem. This is the default. (This option is ignored when noatime is set.)",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodiratime",
		"Do not update directory inode access times on this filesystem. (This option is implied when noatime is set.)",
	),
	docvalues.CreateEnumStringWithDoc(
		"dirsync",
		"All directory updates within the filesystem should be done synchronously. This affects the following system calls: creat(2), link(2), unlink(2), symlink(2), mkdir(2), rmdir(2), mknod(2) and rename(2).",
	),
	docvalues.CreateEnumStringWithDoc(
		"exec",
		"Permit execution of binaries and other executable files.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noexec",
		"Do not permit direct execution of any binaries on the mounted filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"group",
		"Allow an ordinary user to mount the filesystem if one of that user’s groups matches the group of the device. This option implies the options nosuid and nodev (unless overridden by subsequent options, as in the option line group,dev,suid).",
	),
	docvalues.CreateEnumStringWithDoc(
		"iversion",
		"Every time the inode is modified, the i_version field will be incremented.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noiversion",
		"Do not increment the i_version inode field.",
	),
	docvalues.CreateEnumStringWithDoc(
		"mand",
		"Allow mandatory locks on this filesystem. See fcntl(2). This option was deprecated in Linux 5.15.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nomand",
		"Do not allow mandatory locks on this filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"_netdev",
		"The filesystem resides on a device that requires network access (used to prevent the system from attempting to mount these filesystems until the network has been enabled on the system).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nofail",
		"Do not report errors for this device if it does not exist.",
	),
	docvalues.CreateEnumStringWithDoc(
		"relatime",
		"Update inode access times relative to modify or change time. Access time is only updated if the previous access time was earlier than or equal to the current modify or change time. (Similar to noatime, but it doesn’t break mutt(1) or other applications that need to know if a file has been read since the last time it was modified.)",
	),
	docvalues.CreateEnumStringWithDoc(
		"norelatime",
		"Do not use the relatime feature. See also the strictatime mount option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"strictatime",
		"Allows to explicitly request full atime updates. This makes it possible for the kernel to default to relatime or noatime but still allow userspace to override it. For more details about the default system mount options see /proc/mounts.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nostrictatime",
		"Use the kernel’s default behavior for inode access time updates.",
	),
	docvalues.CreateEnumStringWithDoc(
		"lazytime",
		"Only update times (atime, mtime, ctime) on the in-memory version of the file inode. This mount option significantly reduces writes to the inode table for workloads that perform frequent random writes to preallocated files.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nolazytime",
		"Do not use the lazytime feature.",
	),
	docvalues.CreateEnumStringWithDoc(
		"suid",
		"Honor set-user-ID and set-group-ID bits or file capabilities when executing programs from this filesystem.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nosuid",
		"Do not honor set-user-ID and set-group-ID bits or file capabilities when executing programs from this filesystem. In addition, SELinux domain transitions require permission nosuid_transition, which in turn needs also policy capability nnp_nosuid_transition.",
	),
	docvalues.CreateEnumStringWithDoc(
		"silent",
		"Turn on the silent flag.",
	),
	docvalues.CreateEnumStringWithDoc(
		"loud",
		"Turn off the silent flag.",
	),
	docvalues.CreateEnumStringWithDoc(
		"owner",
		"Allow an ordinary user to mount the filesystem if that user is the owner of the device. This option implies the options nosuid and nodev (unless overridden by subsequent options, as in the option line owner,dev,suid).",
	),
	docvalues.CreateEnumStringWithDoc(
		"remount",
		"Attempt to remount an already-mounted filesystem. This is commonly used to change the mount flags for a filesystem, especially to make a readonly filesystem writable. It does not change device or mount point.",
	),
	docvalues.CreateEnumStringWithDoc(
		"ro",
		"Mount the filesystem read-only.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rw",
		"Mount the filesystem read-write.",
	),
	docvalues.CreateEnumStringWithDoc(
		"sync",
		"All I/O to the filesystem should be done synchronously. In the case of media with a limited number of write cycles (e.g. some flash drives), sync may cause life-cycle shortening.",
	),
	docvalues.CreateEnumStringWithDoc(
		"user",
		"Allow an ordinary user to mount the filesystem. The name of the mounting user is written to the mtab file (or to the private libmount file in /run/mount on systems without a regular mtab) so that this same user can unmount the filesystem again. This option implies the options noexec, nosuid, and nodev (unless overridden by subsequent options, as in the option line user,exec,dev,suid).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nouser",
		"Forbid an ordinary user to mount the filesystem. This is the default; it does not imply any other options.",
	),
	docvalues.CreateEnumStringWithDoc(
		"users",
		"Allow any user to mount and to unmount the filesystem, even when some other ordinary user mounted it. This option implies the options noexec, nosuid, and nodev (unless overridden by subsequent options, as in the option line users,exec,dev,suid).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nosymfollow",
		"Do not follow symlinks when resolving paths. Symlinks can still be created, and readlink(1), readlink(2), realpath(1), and realpath(3) all still work properly.",
	),
}

type assignOption struct {
	Documentation string
	Handler       func(context docvalues.KeyValueAssignmentContext) docvalues.Value
}

var defaultAssignOptions = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"context",
		"The context= option is useful when mounting filesystems that do not support extended attributes, such as a floppy or hard disk formatted with VFAT, or systems that are not normally running under SELinux, such as an ext3 or ext4 formatted disk from a non-SELinux workstation. You can also use context= on filesystems you do not trust, such as a floppy. It also helps in compatibility with xattr-supporting filesystems on earlier 2.4.<x> kernel versions. Even where xattrs are supported, you can save time not having to label every file by assigning the entire disk one security context. A commonly used option for removable media is context=\"system_u:object_r:removable_t\".",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"fscontext",
		"The fscontext= option works for all filesystems, regardless of their xattr support. The fscontext option sets the overarching filesystem label to a specific security context. This filesystem label is separate from the individual labels on the files. It represents the entire filesystem for certain kinds of permission checks, such as during mount or file creation. Individual file labels are still obtained from the xattrs on the files themselves. The context option actually sets the aggregate context that fscontext provides, in addition to supplying the same label for individual files.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"defcontext",
		"You can set the default security context for unlabeled files using defcontext= option. This overrides the value set for unlabeled files in the policy and requires a filesystem that supports xattr labeling.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"rootcontext",
		"The rootcontext= option allows you to explicitly label the root inode of a FS being mounted before that FS or inode becomes visible to userspace. This was found to be useful for things like stateless Linux. The special value @target can be used to assign the current context of the target mountpoint location.",
	): docvalues.StringValue{},
}

func createMountOptionField(
	options []docvalues.EnumString,
	assignOption map[docvalues.EnumString]docvalues.Value,
) docvalues.Value {
	dynamicOptions := docvalues.MergeKeyEnumAssignmentMaps(defaultAssignOptions, assignOption)

	return docvalues.ArrayValue{
		Separator:           ",",
		DuplicatesExtractor: &mountOptionsExtractor,
		SubValue: docvalues.OrValue{
			Values: []docvalues.Value{
				docvalues.KeyEnumAssignmentValue{
					Values:          dynamicOptions,
					ValueIsOptional: false,
					Separator:       "=",
				},
				docvalues.EnumValue{
					EnforceValues: true,
					Values:        append(defaultOptions, options...),
				},
			},
		},
	}
}

var DefaultMountOptionsField = createMountOptionField([]docvalues.EnumString{}, map[docvalues.EnumString]docvalues.Value{})

var MountOptionsMapField = map[string]docvalues.Value{
	"adfs": createMountOptionField(
		commondocumentation.AdfsDocumentationEnums,
		commondocumentation.AdfsDocumentationAssignable,
	),
	"affs": createMountOptionField(
		commondocumentation.AffsDocumentationEnums,
		commondocumentation.AffsDocumentationAssignable,
	),
	"debugfs": createMountOptionField(
		commondocumentation.DebugfsDocumentationEnums,
		commondocumentation.DebugfsDocumentationAssignable,
	),
	"ext2": createMountOptionField(
		commondocumentation.Ext2DocumentationEnums,
		commondocumentation.Ext2DocumentationAssignable,
	),
	"ext3": createMountOptionField(
		append(commondocumentation.Ext2DocumentationEnums, commondocumentation.Ext3DocumentationEnums...),
		docvalues.MergeKeyEnumAssignmentMaps(commondocumentation.Ext2DocumentationAssignable, commondocumentation.Ext3DocumentationAssignable),
	),
	"ext4": createMountOptionField(
		append(append(commondocumentation.Ext2DocumentationEnums, commondocumentation.Ext3DocumentationEnums...), commondocumentation.Ext4DocumentationEnums...),
		docvalues.MergeKeyEnumAssignmentMaps(commondocumentation.Ext2DocumentationAssignable, docvalues.MergeKeyEnumAssignmentMaps(commondocumentation.Ext3DocumentationAssignable, commondocumentation.Ext4DocumentationAssignable)),
	),
	"devpts": createMountOptionField(
		commondocumentation.DevptsDocumentationEnums,
		commondocumentation.DevptsDocumentationAssignable,
	),
	"fat": createMountOptionField(
		commondocumentation.FatDocumentationEnums,
		commondocumentation.FatDocumentationAssignable,
	),
	"hfs": createMountOptionField(
		commondocumentation.HfsDocumentationEnums,
		commondocumentation.HfsDocumentationAssignable,
	),
	"hpfs": createMountOptionField(
		commondocumentation.HpfsDocumentationEnums,
		commondocumentation.HpfsDocumentationAssignable,
	),
}
