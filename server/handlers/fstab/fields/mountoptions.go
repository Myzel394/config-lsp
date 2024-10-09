package fields

import (
	commondocumentation "config-lsp/common-documentation"
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
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.automount",
		`An automount unit will be created for the file system. See systemd.automount(5) for details.

Added in version 215.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.makefs",
		`The file system will be initialized on the device. If the device is not "empty", i.e. it contains any signature, the operation will be skipped. It is hence expected that this option remains set even after the device has been initialized.

Note that this option can only be used in /etc/fstab, and will be ignored when part of the Options= setting in a unit file.

See systemd-makefs@.service(8).

wipefs(8) may be used to remove any signatures from a block device to force x-systemd.makefs to reinitialize the device.

Added in version 236.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.growfs",
		`The file system will be grown to occupy the full block device. If the file system is already at maximum size, no action will be performed. It is hence expected that this option remains set even after the file system has been grown. Only certain file system types are supported, see systemd-makefs@.service(8) for details.

Note that this option can only be used in /etc/fstab, and will be ignored when part of the Options= setting in a unit file.

Added in version 236.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.pcrfs",
		`Measures file system identity information (mount point, type, label, UUID, partition label, partition UUID) into PCR 15 after the file system has been mounted. This ensures the systemd-pcrfs@.service(8) or systemd-pcrfs-root.service services are pulled in by the mount unit.

Note that this option can only be used in /etc/fstab, and will be ignored when part of the Options= setting in a unit file. It is also implied for the root and /usr/ partitions discovered by systemd-gpt-auto-generator(8).

Added in version 253.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.rw-only",
		`If a mount operation fails to mount the file system read-write, it normally tries mounting the file system read-only instead. This option disables that behaviour, and causes the mount to fail immediately instead. This option is translated into the ReadWriteOnly= setting in a unit file.

Added in version 246.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"x-initrd.mount",
		`An additional filesystem to be mounted in the initrd. See initrd-fs.target description in systemd.special(7). This is both an indicator to the initrd to mount this partition early and an indicator to the host to leave the partition mounted until final shutdown. Or in other words, if this flag is set it is assumed the mount shall be active during the entire regular runtime of the system, i.e. established before the initrd transitions into the host all the way until the host transitions to the final shutdown phase.

Added in version 215.`,
	),
}

type assignOption struct {
	Documentation string
	Handler       func(context docvalues.KeyValueAssignmentContext) docvalues.DeprecatedValue
}

var defaultAssignOptions = map[docvalues.EnumString]docvalues.DeprecatedValue{
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
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.requires",
		`Configures a Requires= and an After= dependency between the created mount unit and another systemd unit, such as a device or mount unit. The argument should be a unit name, or an absolute path to a device node or mount point. This option may be specified more than once. This option is particularly useful for mount point declarations that need an additional device to be around (such as an external journal device for journal file systems) or an additional mount to be in place (such as an overlay file system that merges multiple mount points). See After= and Requires= in systemd.unit(5) for details.

Note that this option always applies to the created mount unit only regardless whether x-systemd.automount has been specified.

Added in version 220.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.before",
		`In the created mount unit, configures a Before= or After= dependency on another systemd unit, such as a mount unit. The argument should be a unit name or an absolute path to a mount point. This option may be specified more than once. This option is particularly useful for mount point declarations with nofail option that are mounted asynchronously but need to be mounted before or after some unit start, for example, before local-fs.target unit. See Before= and After= in systemd.unit(5) for details.

Note that these options always apply to the created mount unit only regardless whether x-systemd.automount has been specified.

Added in version 233.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.after",
		`In the created mount unit, configures a Before= or After= dependency on another systemd unit, such as a mount unit. The argument should be a unit name or an absolute path to a mount point. This option may be specified more than once. This option is particularly useful for mount point declarations with nofail option that are mounted asynchronously but need to be mounted before or after some unit start, for example, before local-fs.target unit. See Before= and After= in systemd.unit(5) for details.

Note that these options always apply to the created mount unit only regardless whether x-systemd.automount has been specified.

Added in version 233.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.wanted-by",
		`In the created mount unit, configures a WantedBy= or RequiredBy= dependency on another unit. This option may be specified more than once. If this is specified, the default dependencies (see above) other than umount.target on the created mount unit, e.g.  local-fs.target, are not automatically created. Hence it is likely that some ordering dependencies need to be set up manually through x-systemd.before= and x-systemd.after=. See WantedBy= and RequiredBy= in systemd.unit(5) for details.

Added in version 245.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.required-by",
		`In the created mount unit, configures a WantedBy= or RequiredBy= dependency on another unit. This option may be specified more than once. If this is specified, the default dependencies (see above) other than umount.target on the created mount unit, e.g.  local-fs.target, are not automatically created. Hence it is likely that some ordering dependencies need to be set up manually through x-systemd.before= and x-systemd.after=. See WantedBy= and RequiredBy= in systemd.unit(5) for details.

Added in version 245.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.wants-mounts-for",
		`Configures a RequiresMountsFor= or WantsMountsFor= dependency between the created mount unit and other mount units. The argument must be an absolute path. This option may be specified more than once. See RequiresMountsFor= or WantsMountsFor= in systemd.unit(5) for details.

Added in version 220.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.requires-mounts-for",
		`Configures a RequiresMountsFor= or WantsMountsFor= dependency between the created mount unit and other mount units. The argument must be an absolute path. This option may be specified more than once. See RequiresMountsFor= or WantsMountsFor= in systemd.unit(5) for details.

Added in version 220.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.device-bound",
		`Takes a boolean argument. If true or no argument, a BindsTo= dependency on the backing device is set. If false, the mount unit is not stopped no matter whether the backing device is still present. This is useful when the file system is backed by volume managers. If not set, and the mount comes from unit fragments, i.e. generated from /etc/fstab by systemd-fstab-generator(8) or loaded from a manually configured mount unit, a combination of Requires= and StopPropagatedFrom= dependencies is set on the backing device. If doesn't, only Requires= is used.

Added in version 233.`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("true"),
			docvalues.CreateEnumString("false"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.idle-timeout",
		`Configures the idle timeout of the automount unit. See TimeoutIdleSec= in systemd.automount(5) for details.

Added in version 220.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.device-timeout",
		`Configure how long systemd should wait for a device to show up before giving up on an entry from /etc/fstab. Specify a time in seconds or explicitly append a unit such as "s", "min", "h", "ms".

Note that this option can only be used in /etc/fstab, and will be ignored when part of the Options= setting in a unit file.

Added in version 215.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"x-systemd.mount-timeout",
		`Configure how long systemd should wait for the mount command to finish before giving up on an entry from /etc/fstab. Specify a time in seconds or explicitly append a unit such as "s", "min", "h", "ms".

Note that this option can only be used in /etc/fstab, and will be ignored when part of the Options= setting in a unit file.

See TimeoutSec= below for details.

Added in version 233.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"fscontext",
		"The fscontext= option works for all filesystems, regardless of their xattr support. The fscontext option sets the overarching filesystem label to a specific security context. This filesystem label is separate from the individual labels on the files. It represents the entire filesystem for certain kinds of permission checks, such as during mount or file creation. Individual file labels are still obtained from the xattrs on the files themselves. The context option actually sets the aggregate context that fscontext provides, in addition to supplying the same label for individual files.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"defcontext",
		"You can set the default security context for unlabeled files using defcontext= option. This overrides the value set for unlabeled files in the policy and requires a filesystem that supports xattr labeling.",
	): docvalues.StringValue{},
}

func createMountOptionField(
	options []docvalues.EnumString,
	assignOption map[docvalues.EnumString]docvalues.DeprecatedValue,
) docvalues.DeprecatedValue {
	// dynamicOptions := docvalues.MergeKeyEnumAssignmentMaps(defaultAssignOptions, assignOption)

	return docvalues.ArrayValue{
		Separator:           ",",
		DuplicatesExtractor: &mountOptionsExtractor,
		SubValue: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
				docvalues.KeyEnumAssignmentValue{
					Values:          assignOption,
					ValueIsOptional: false,
					Separator:       "=",
				},
				docvalues.EnumValue{
					EnforceValues: true,
					Values:        options,
				},
			},
		},
	}
}

var DefaultMountOptionsField = createMountOptionField(defaultOptions, defaultAssignOptions)

var MountOptionsMapField = map[string]docvalues.DeprecatedValue{
	"adfs": createMountOptionField(
		commondocumentation.AdfsDocumentationEnums,
		commondocumentation.AdfsDocumentationAssignable,
	),
	"affs": createMountOptionField(
		commondocumentation.AffsDocumentationEnums,
		commondocumentation.AffsDocumentationAssignable,
	),
	"btrfs": createMountOptionField(
		commondocumentation.BtrfsDocumentationEnums,
		commondocumentation.BtrfsDocumentationAssignable,
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
	"iso9660": createMountOptionField(
		commondocumentation.Iso9660DocumentationEnums,
		commondocumentation.Iso9660DocumentationAssignable,
	),
	"jfs": createMountOptionField(
		commondocumentation.JfsDocumentationEnums,
		commondocumentation.JfsDocumentationAssignable,
	),
	"msdos": createMountOptionField(
		commondocumentation.MsdosDocumentationEnums,
		commondocumentation.MsdosDocumentationAssignable,
	),
	"ncpfs": createMountOptionField(
		commondocumentation.NcpfsDocumentationEnums,
		commondocumentation.NcpfsDocumentationAssignable,
	),
	"ntfs": createMountOptionField(
		commondocumentation.NtfsDocumentationEnums,
		commondocumentation.NtfsDocumentationAssignable,
	),
	"overlay": createMountOptionField(
		commondocumentation.OverlayDocumentationEnums,
		commondocumentation.OverlayDocumentationAssignable,
	),
	"reiserfs": createMountOptionField(
		commondocumentation.ReiserfsDocumentationEnums,
		commondocumentation.ReiserfsDocumentationAssignable,
	),
	"usbfs": createMountOptionField(
		commondocumentation.UsbfsDocumentationEnums,
		commondocumentation.UsbfsDocumentationAssignable,
	),
	"ubifs": createMountOptionField(
		commondocumentation.UbifsDocumentationEnums,
		commondocumentation.UbifsDocumentationAssignable,
	),
	"udf": createMountOptionField(
		commondocumentation.UdfDocumentationEnums,
		commondocumentation.UdfDocumentationAssignable,
	),
	"ufs": createMountOptionField(
		commondocumentation.UfsDocumentationEnums,
		commondocumentation.UfsDocumentationAssignable,
	),
	"umsdos": createMountOptionField(
		commondocumentation.UmsdosDocumentationEnums,
		commondocumentation.UmsdosDocumentationAssignable,
	),
	"vfat": createMountOptionField(
		commondocumentation.VfatDocumentationEnums,
		commondocumentation.VfatDocumentationAssignable,
	),
}
