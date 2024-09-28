package commondocumentation

import docvalues "config-lsp/doc-values"

var maxInlineMin = 2048

var BtrfsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"check_int_print_mask",
		"These debugging options control the behavior of the integrity checking module (the BTRFS_FS_CHECK_INTEGRITY config option required). The main goal is to verify that all blocks from a given transaction period are properly linked.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"commit",
		"Set the interval of periodic transaction commit when data are synchronized to permanent storage. Higher interval values lead to larger amount of unwritten data, which has obvious consequences when the system crashes. The upper bound is not forced, but a warning is printed if it's more than 300 seconds (5 minutes). Use with care.",
	): docvalues.NumberValue{},
	docvalues.CreateEnumStringWithDoc(
		"compress",
		"Control BTRFS file data compression. Type may be specified as zlib, lzo, zstd or no (for no compression, used for remounting). If no type is specified, zlib is used. If compress-force is specified, then compression will always be attempted, but the data may end up uncompressed if the compression would make them larger.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("zlib"),
			docvalues.CreateEnumString("lzo"),
			docvalues.CreateEnumString("zstd"),
			docvalues.CreateEnumStringWithDoc(
				"no",
				"No compression, used for remounting.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"compress-force",
		"Control BTRFS file data compression. Type may be specified as zlib, lzo, zstd or no (for no compression, used for remounting). If no type is specified, zlib is used. If compress-force is specified, then compression will always be attempted, but the data may end up uncompressed if the compression would make them larger.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("zlib"),
			docvalues.CreateEnumString("lzo"),
			docvalues.CreateEnumString("zstd"),
			docvalues.CreateEnumStringWithDoc(
				"no",
				"No compression, used for remounting.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"device",
		"Specify a path to a device that will be scanned for BTRFS filesystem during mount. This is usually done automatically by a device manager (like udev) or using the btrfs device scan command (eg. run from the initial ramdisk). In cases where this is not possible the device mount option can help.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"fatal_errors",
		"Action to take when encountering a fatal error.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"bug",
				"BUG() on a fatal error, the system will stay in the crashed state and may be still partially usable, but reboot is required for full operation",
			),
			docvalues.CreateEnumStringWithDoc(
				"panic",
				"panic() on a fatal error, depending on other system configuration, this may be followed by a reboot. Please refer to the documentation of kernel boot parameters, e.g. panic, oops or crashkernel.",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"fragment",
		"A debugging helper to intentionally fragment given type of block groups. The type can be data, metadata or all. This mount option should not be used outside of debugging environments and is not recognized if the kernel config option BTRFS_DEBUG is not enabled.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"max_inline",
		"Specify the maximum amount of space, that can be inlined in a metadata b-tree leaf. The value is specified in bytes, optionally with a K suffix (case insensitive). In practice, this value is limited by the filesystem block size (named sectorsize at mkfs time), and memory page size of the system. In case of sectorsize limit, there's some space unavailable due to leaf headers. For example, a 4Ki",
	): docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.EnumValue{
				EnforceValues: true,
				Values: []docvalues.EnumString{
					docvalues.CreateEnumStringWithDoc(
						"0",
						"panic() on a fatal error, depending on other system configuration, this may be followed by a reboot. Please refer to the documentation of kernel boot parameters, e.g. panic, oops or crashkernel.",
					),
				},
			},
			docvalues.SuffixWithMeaningValue{
				Suffixes: []docvalues.Suffix{
					{
						Suffix:  "K",
						Meaning: "Kilobytes",
					},
				},
				SubValue: docvalues.NumberValue{Min: &maxInlineMin},
			},
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"metadata_ratio",
		"Specifies that 1 metadata chunk should be allocated after every value data chunks. Default behaviour depends on internal logic, some percent of unused metadata space is attempted to be maintained but is not always possible if there's not enough space left for chunk allocation. The option could be useful to override the internal logic in favor of the metadata allocation if the expected workload",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"subvol",
		"Mount subvolume from path rather than the toplevel subvolume. The path is always treated as relative to the toplevel subvolume. This mount option overrides the default subvolume set for the given filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"subvolid",
		"Mount subvolume specified by a subvolid number rather than the toplevel subvolume. You can use btrfs subvolume list of btrfs subvolume show to see subvolume ID numbers. This mount option overrides the default subvolume set for the given filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"thread_pool",
		"The number of worker threads to start. NRCPUS is number of on-line CPUs detected at the time of mount. Small number leads to less parallelism in processing data and metadata, higher numbers could lead to a performance hit due to increased locking contention, process scheduling, cache-line bouncing or costly data transfers between local CPU memories.",
	): docvalues.PositiveNumberValue(),
}

var BtrfsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"acl",
		"Enable/disable support for Posix Access Control Lists (ACLs). See the [acl(5)](https://manpages.debian.org/testing/acl/acl.5.en.html) manual page for more information about ACLs.\nThe support for ACL is build-time configurable (BTRFS_FS_POSIX_ACL) and mount fails if acl is requested but the feature is not compiled in.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noacl",
		"Enable/disable support for Posix Access Control Lists (ACLs). See the [acl(5)](https://manpages.debian.org/testing/acl/acl.5.en.html) manual page for more information about ACLs.\nThe support for ACL is build-time configurable (BTRFS_FS_POSIX_ACL) and mount fails if acl is requested but the feature is not compiled in.",
	),
	docvalues.CreateEnumStringWithDoc(
		"autodefrag",
		"Enable automatic file defragmentation. When enabled, small random writes into files (in a range of tens of kilobytes, currently it's 64KiB) are detected and queued up for the defragmentation process. Not well suited for large database workloads.\nThe read latency may increase due to reading the adjacent blocks that make up the range for defragmentation, successive write will merge the blocks in the new location.\nWARNING:\nDefragmenting with Linux kernel versions < 3.9 or ≥ 3.14-rc2 as well as with Linux stable kernel versions ≥ 3.10.31, ≥ 3.12.12 or ≥ 3.13.4 will break up the reflinks of COW data (for example files copied with cp --reflink, snapshots or de-duplicated data). This may cause considerable increase of space usage depending on the broken up reflinks.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noautodefrag",
		"Enable automatic file defragmentation. When enabled, small random writes into files (in a range of tens of kilobytes, currently it's 64KiB) are detected and queued up for the defragmentation process. Not well suited for large database workloads.\nThe read latency may increase due to reading the adjacent blocks that make up the range for defragmentation, successive write will merge the blocks in the new location.\nWARNING:\nDefragmenting with Linux kernel versions < 3.9 or ≥ 3.14-rc2 as well as with Linux stable kernel versions ≥ 3.10.31, ≥ 3.12.12 or ≥ 3.13.4 will break up the reflinks of COW data (for example files copied with cp --reflink, snapshots or de-duplicated data). This may cause considerable increase of space usage depending on the broken up reflinks.",
	),
	docvalues.CreateEnumStringWithDoc(
		"barrier",
		"Ensure that all IO write operations make it through the device cache and are stored permanently when the filesystem is at its consistency checkpoint. This typically means that a flush command is sent to the device that will synchronize all pending data and ordinary metadata blocks, then writes the superblock and issues another flush.\nThe write flushes incur a slight hit and also prevent the IO block scheduler to reorder requests in a more effective way. Disabling barriers gets rid of that penalty but will most certainly lead to a corrupted filesystem in case of a crash or power loss. The ordinary metadata blocks could be yet unwritten at the time the new superblock is stored permanently, expecting that the block pointers t\nOn a device with a volatile battery-backed write-back cache, the nobarrier option will not lead to filesystem corruption as the pending blocks are supposed to make it to the permanent storage.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nobarrier",
		"Ensure that all IO write operations make it through the device cache and are stored permanently when the filesystem is at its consistency checkpoint. This typically means that a flush command is sent to the device that will synchronize all pending data and ordinary metadata blocks, then writes the superblock and issues another flush.\nThe write flushes incur a slight hit and also prevent the IO block scheduler to reorder requests in a more effective way. Disabling barriers gets rid of that penalty but will most certainly lead to a corrupted filesystem in case of a crash or power loss. The ordinary metadata blocks could be yet unwritten at the time the new superblock is stored permanently, expecting that the block pointers t\nOn a device with a volatile battery-backed write-back cache, the nobarrier option will not lead to filesystem corruption as the pending blocks are supposed to make it to the permanent storage.",
	),
	docvalues.CreateEnumStringWithDoc(
		"check_int",
		"These debugging options control the behavior of the integrity checking module (the BTRFS_FS_CHECK_INTEGRITY config option required). The main goal is to verify that all blocks from a given transaction period are properly linked.\ncheck_int enables the integrity checker module, which examines all block write requests to ensure on-disk consistency, at a large memory and CPU cost.",
	),
	docvalues.CreateEnumStringWithDoc(
		"check_int_data",
		"These debugging options control the behavior of the integrity checking module (the BTRFS_FS_CHECK_INTEGRITY config option required). The main goal is to verify that all blocks from a given transaction period are properly linked.\ncheck_int_data includes extent data in the integrity checks, and implies the check_int option.",
	),
	docvalues.CreateEnumStringWithDoc(
		"clear_cache",
		"Force clearing and rebuilding of the disk space cache if something has gone wrong. See also: space_cache.",
	),
	docvalues.CreateEnumStringWithDoc(
		"datacow",
		"Enable data copy-on-write for newly created files. Nodatacow implies nodatasum, and disables compression. All files created under nodatacow are also set the NOCOW file attribute (see [chattr(1)](https://manpages.debian.org/testing/e2fsprogs/chattr.1.en.html)).\nNOTE:\nIf nodatacow or nodatasum are enabled, compression is disabled.\nUpdates in-place improve performance for workloads that do frequent overwrites, at the cost of potential partial writes, in case the write is interrupted (system crash, device failure).",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodatacow",
		"Enable data copy-on-write for newly created files. Nodatacow implies nodatasum, and disables compression. All files created under nodatacow are also set the NOCOW file attribute (see [chattr(1)](https://manpages.debian.org/testing/e2fsprogs/chattr.1.en.html)).\nNOTE:\nIf nodatacow or nodatasum are enabled, compression is disabled.\nUpdates in-place improve performance for workloads that do frequent overwrites, at the cost of potential partial writes, in case the write is interrupted (system crash, device failure).",
	),
	docvalues.CreateEnumStringWithDoc(
		"datasum",
		"Enable data checksumming for newly created files. Datasum implies datacow, ie. the normal mode of operation. All files created under nodatasum inherit the \"no checksums\" property, however there's no corresponding file attribute (see [chattr(1)](https://manpages.debian.org/testing/e2fsprogs/chattr.1.en.html)).\nNOTE:\nIf nodatacow or nodatasum are enabled, compression is disabled.\nThere is a slight performance gain when checksums are turned off, the corresponding metadata blocks holding the checksums do not need to updated. The cost of checksumming of the blocks in memory is much lower than the IO, modern CPUs feature hardware support of the checksumming algorithm.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodatasum",
		"Enable data checksumming for newly created files. Datasum implies datacow, ie. the normal mode of operation. All files created under nodatasum inherit the \"no checksums\" property, however there's no corresponding file attribute (see [chattr(1)](https://manpages.debian.org/testing/e2fsprogs/chattr.1.en.html)).\nNOTE:\nIf nodatacow or nodatasum are enabled, compression is disabled.\nThere is a slight performance gain when checksums are turned off, the corresponding metadata blocks holding the checksums do not need to updated. The cost of checksumming of the blocks in memory is much lower than the IO, modern CPUs feature hardware support of the checksumming algorithm.",
	),
	docvalues.CreateEnumStringWithDoc(
		"degraded",
		"Allow mounts with less devices than the RAID profile constraints require. A read-write mount (or remount) may fail when there are too many devices missing, for example if a stripe member is completely missing from RAID0.\nSince 4.14, the constraint checks have been improved and are verified on the chunk level, not an the device level. This allows degraded mounts of filesystems with mixed RAID profiles for data and metadata, even if the device number constraints would not be satisfied for some of the profiles.\nExample: metadata -- raid1, data -- single, devices -- /dev/sda, /dev/sdb\nSuppose the data are completely stored on sda, then missing sdb will not prevent the mount, even if 1 missing device would normally prevent (any) single profile to mount. In case some of the data chunks are stored on sdb, then the constraint of single/data is not satisfied and the filesystem cannot be mounted.",
	),
	docvalues.CreateEnumStringWithDoc(
		"discard",
		"Enable discarding of freed file blocks. This is useful for SSD devices, thinly provisioned LUNs, or virtual machine images; however, every storage layer must support discard for it to work.\nIn the synchronous mode (sync or without option value), lack of asynchronous queued TRIM on the backing device TRIM can severely degrade performance, because a synchronous TRIM operation will be attempted instead. Queued TRIM requires newer than SATA revision 3.1 chipsets and devices.\nThe asynchronous mode (async) gathers extents in larger chunks before sending them to the devices for TRIM. The overhead and performance impact should be negligible compared to the previous mode and it's supposed to be the preferred mode if needed.\nIf it is not necessary to immediately discard freed blocks, then the fstrim tool can be used to discard all free blocks in a batch. Scheduling a TRIM during a period of low system activity will prevent latent interference with the performance of other operations. Also, a device may ignore the TRIM command if the range is too small, so running a batch discard has a greater probability of actual",
	),
	docvalues.CreateEnumStringWithDoc(
		"discard=sync",
		"Enable discarding of freed file blocks. This is useful for SSD devices, thinly provisioned LUNs, or virtual machine images; however, every storage layer must support discard for it to work.\nIn the synchronous mode (sync or without option value), lack of asynchronous queued TRIM on the backing device TRIM can severely degrade performance, because a synchronous TRIM operation will be attempted instead. Queued TRIM requires newer than SATA revision 3.1 chipsets and devices.\nThe asynchronous mode (async) gathers extents in larger chunks before sending them to the devices for TRIM. The overhead and performance impact should be negligible compared to the previous mode and it's supposed to be the preferred mode if needed.\nIf it is not necessary to immediately discard freed blocks, then the fstrim tool can be used to discard all free blocks in a batch. Scheduling a TRIM during a period of low system activity will prevent latent interference with the performance of other operations. Also, a device may ignore the TRIM command if the range is too small, so running a batch discard has a greater probability of actual",
	),
	docvalues.CreateEnumStringWithDoc(
		"discard=async",
		"Enable discarding of freed file blocks. This is useful for SSD devices, thinly provisioned LUNs, or virtual machine images; however, every storage layer must support discard for it to work.\nIn the synchronous mode (sync or without option value), lack of asynchronous queued TRIM on the backing device TRIM can severely degrade performance, because a synchronous TRIM operation will be attempted instead. Queued TRIM requires newer than SATA revision 3.1 chipsets and devices.\nThe asynchronous mode (async) gathers extents in larger chunks before sending them to the devices for TRIM. The overhead and performance impact should be negligible compared to the previous mode and it's supposed to be the preferred mode if needed.\nIf it is not necessary to immediately discard freed blocks, then the fstrim tool can be used to discard all free blocks in a batch. Scheduling a TRIM during a period of low system activity will prevent latent interference with the performance of other operations. Also, a device may ignore the TRIM command if the range is too small, so running a batch discard has a greater probability of actual",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodiscard",
		"Enable discarding of freed file blocks. This is useful for SSD devices, thinly provisioned LUNs, or virtual machine images; however, every storage layer must support discard for it to work.\nIn the synchronous mode (sync or without option value), lack of asynchronous queued TRIM on the backing device TRIM can severely degrade performance, because a synchronous TRIM operation will be attempted instead. Queued TRIM requires newer than SATA revision 3.1 chipsets and devices.\nThe asynchronous mode (async) gathers extents in larger chunks before sending them to the devices for TRIM. The overhead and performance impact should be negligible compared to the previous mode and it's supposed to be the preferred mode if needed.\nIf it is not necessary to immediately discard freed blocks, then the fstrim tool can be used to discard all free blocks in a batch. Scheduling a TRIM during a period of low system activity will prevent latent interference with the performance of other operations. Also, a device may ignore the TRIM command if the range is too small, so running a batch discard has a greater probability of actual",
	),
	docvalues.CreateEnumStringWithDoc(
		"enospc_debug",
		"Enable verbose output for some ENOSPC conditions. It's safe to use but can be noisy if the system reaches near-full state.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noenospc_debug",
		"Enable verbose output for some ENOSPC conditions. It's safe to use but can be noisy if the system reaches near-full state.",
	),
	docvalues.CreateEnumStringWithDoc(
		"flushoncommit",
		"This option forces any data dirtied by a write in a prior transaction to commit as part of the current commit, effectively a full filesystem sync.\nThis makes the committed state a fully consistent view of the file system from the application's perspective (i.e. it includes all completed file system operations). This was previously the behavior only when a snapshot was created.\nWhen off, the filesystem is consistent but buffered writes may last more than one transaction commit.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noflushoncommit",
		"This option forces any data dirtied by a write in a prior transaction to commit as part of the current commit, effectively a full filesystem sync.\nThis makes the committed state a fully consistent view of the file system from the application's perspective (i.e. it includes all completed file system operations). This was previously the behavior only when a snapshot was created.\nWhen off, the filesystem is consistent but buffered writes may last more than one transaction commit.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nologreplay",
		"The tree-log contains pending updates to the filesystem until the full commit. The log is replayed on next mount, this can be disabled by this option. See also treelog. Note that nologreplay is the same as norecovery.\nWARNING:\nCurrently, the tree log is replayed even with a read-only mount! To disable that behaviour, mount also with nologreplay.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rescan_uuid_tree",
		"Force check and rebuild procedure of the UUID tree. This should not normally be needed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"rescue",
		"Modes allowing mount with damaged filesystem structures.\nusebackuproot (since: 5.9, replaces standalone option usebackuproot)\nnologreplay (since: 5.9, replaces standalone option nologreplay)\nignorebadroots, ibadroots (since: 5.11)\nignoredatacsums, idatacsums (since: 5.11)\nall (since: 5.9)",
	),
	docvalues.CreateEnumStringWithDoc(
		"skip_balance",
		"Skip automatic resume of an interrupted balance operation. The operation can later be resumed with btrfs balance resume, or the paused state can be removed with btrfs balance cancel. The default behaviour is to resume an interrupted balance immediately after a volume is mounted.",
	),
	docvalues.CreateEnumStringWithDoc(
		"space_cache",
		"Options to control the free space cache. The free space cache greatly improves performance when reading block group free space into memory. However, managing the space cache consumes some resources, including a small amount of disk space.\nThere are two implementations of the free space cache. The original one, referred to as v1, is the safe default. The v1 space cache can be disabled at mount time with nospace_cache without clearing.\nOn very large filesystems (many terabytes) and certain workloads, the performance of the v1 space cache may degrade drastically. The v2 implementation, which adds a new b-tree called the free space tree, addresses this issue. Once enabled, the v2 space cache will always be used and cannot be disabled unless it is cleared. Use clear_cache,space_cache=v1 or clear_cache,nospace_cache to do so. If\nThe [btrfs-check(8)](https://manpages.debian.org/testing/btrfs-progs/btrfs-check.8.en.html) and `[mkfs.btrfs(8)](https://manpages.debian.org/testing/btrfs-progs/mkfs.btrfs.8.en.html) commands have full v2 free space cache support since v4.19.\nIf a version is not explicitly specified, the default implementation will be chosen, which is v1.",
	),
	docvalues.CreateEnumStringWithDoc(
		"space_cache=v1",
		"Options to control the free space cache. The free space cache greatly improves performance when reading block group free space into memory. However, managing the space cache consumes some resources, including a small amount of disk space.\nThere are two implementations of the free space cache. The original one, referred to as v1, is the safe default. The v1 space cache can be disabled at mount time with nospace_cache without clearing.\nOn very large filesystems (many terabytes) and certain workloads, the performance of the v1 space cache may degrade drastically. The v2 implementation, which adds a new b-tree called the free space tree, addresses this issue. Once enabled, the v2 space cache will always be used and cannot be disabled unless it is cleared. Use clear_cache,space_cache=v1 or clear_cache,nospace_cache to do so. If\nThe [btrfs-check(8)](https://manpages.debian.org/testing/btrfs-progs/btrfs-check.8.en.html) and `[mkfs.btrfs(8)](https://manpages.debian.org/testing/btrfs-progs/mkfs.btrfs.8.en.html) commands have full v2 free space cache support since v4.19.\nIf a version is not explicitly specified, the default implementation will be chosen, which is v1.",
	),
	docvalues.CreateEnumStringWithDoc(
		"space_cache=v2",
		"Options to control the free space cache. The free space cache greatly improves performance when reading block group free space into memory. However, managing the space cache consumes some resources, including a small amount of disk space.\nThere are two implementations of the free space cache. The original one, referred to as v1, is the safe default. The v1 space cache can be disabled at mount time with nospace_cache without clearing.\nOn very large filesystems (many terabytes) and certain workloads, the performance of the v1 space cache may degrade drastically. The v2 implementation, which adds a new b-tree called the free space tree, addresses this issue. Once enabled, the v2 space cache will always be used and cannot be disabled unless it is cleared. Use clear_cache,space_cache=v1 or clear_cache,nospace_cache to do so. If\nThe [btrfs-check(8)](https://manpages.debian.org/testing/btrfs-progs/btrfs-check.8.en.html) and `[mkfs.btrfs(8)](https://manpages.debian.org/testing/btrfs-progs/mkfs.btrfs.8.en.html) commands have full v2 free space cache support since v4.19.\nIf a version is not explicitly specified, the default implementation will be chosen, which is v1.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nospace_cache",
		"Options to control the free space cache. The free space cache greatly improves performance when reading block group free space into memory. However, managing the space cache consumes some resources, including a small amount of disk space.\nThere are two implementations of the free space cache. The original one, referred to as v1, is the safe default. The v1 space cache can be disabled at mount time with nospace_cache without clearing.\nOn very large filesystems (many terabytes) and certain workloads, the performance of the v1 space cache may degrade drastically. The v2 implementation, which adds a new b-tree called the free space tree, addresses this issue. Once enabled, the v2 space cache will always be used and cannot be disabled unless it is cleared. Use clear_cache,space_cache=v1 or clear_cache,nospace_cache to do so. If\nThe [btrfs-check(8)](https://manpages.debian.org/testing/btrfs-progs/btrfs-check.8.en.html) and `[mkfs.btrfs(8)](https://manpages.debian.org/testing/btrfs-progs/mkfs.btrfs.8.en.html) commands have full v2 free space cache support since v4.19.\nIf a version is not explicitly specified, the default implementation will be chosen, which is v1.",
	),
	docvalues.CreateEnumStringWithDoc(
		"ssd",
		"Options to control SSD allocation schemes. By default, BTRFS will enable or disable SSD optimizations depending on status of a device with respect to rotational or non-rotational type. This is determined by the contents of /sys/block/DEV/queue/rotational). If it is 0, the ssd option is turned on. The option nossd will disable the autodetection.\nThe optimizations make use of the absence of the seek penalty that's inherent for the rotational devices. The blocks can be typically written faster and are not offloaded to separate threads.\nNOTE:\nSince 4.14, the block layout optimizations have been dropped. This used to help with first generations of SSD devices. Their FTL (flash translation layer) was not effective and the optimization was supposed to improve the wear by better aligning blocks. This is no longer true with modern SSD devices and the optimization had no real benefit. Furthermore it caused increased fragmentation. The layout",
	),
	docvalues.CreateEnumStringWithDoc(
		"ssd_spread",
		"Options to control SSD allocation schemes. By default, BTRFS will enable or disable SSD optimizations depending on status of a device with respect to rotational or non-rotational type. This is determined by the contents of /sys/block/DEV/queue/rotational). If it is 0, the ssd option is turned on. The option nossd will disable the autodetection.\nThe optimizations make use of the absence of the seek penalty that's inherent for the rotational devices. The blocks can be typically written faster and are not offloaded to separate threads.\nNOTE:\nSince 4.14, the block layout optimizations have been dropped. This used to help with first generations of SSD devices. Their FTL (flash translation layer) was not effective and the optimization was supposed to improve the wear by better aligning blocks. This is no longer true with modern SSD devices and the optimization had no real benefit. Furthermore it caused increased fragmentation. The layout\nThe ssd_spread mount option attempts to allocate into bigger and aligned chunks of unused space, and may perform better on low-end SSDs. ssd_spread implies ssd, enabling all other SSD heuristics as well. The option nossd will disable all SSD options while nossd_spread only disables ssd_spread.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nossd",
		"Options to control SSD allocation schemes. By default, BTRFS will enable or disable SSD optimizations depending on status of a device with respect to rotational or non-rotational type. This is determined by the contents of /sys/block/DEV/queue/rotational). If it is 0, the ssd option is turned on. The option nossd will disable the autodetection.\nThe optimizations make use of the absence of the seek penalty that's inherent for the rotational devices. The blocks can be typically written faster and are not offloaded to separate threads.\nNOTE:\nSince 4.14, the block layout optimizations have been dropped. This used to help with first generations of SSD devices. Their FTL (flash translation layer) was not effective and the optimization was supposed to improve the wear by better aligning blocks. This is no longer true with modern SSD devices and the optimization had no real benefit. Furthermore it caused increased fragmentation. The layout",
	),
	docvalues.CreateEnumStringWithDoc(
		"nossd_spread",
		"Options to control SSD allocation schemes. By default, BTRFS will enable or disable SSD optimizations depending on status of a device with respect to rotational or non-rotational type. This is determined by the contents of /sys/block/DEV/queue/rotational). If it is 0, the ssd option is turned on. The option nossd will disable the autodetection.\nThe optimizations make use of the absence of the seek penalty that's inherent for the rotational devices. The blocks can be typically written faster and are not offloaded to separate threads.\nNOTE:\nSince 4.14, the block layout optimizations have been dropped. This used to help with first generations of SSD devices. Their FTL (flash translation layer) was not effective and the optimization was supposed to improve the wear by better aligning blocks. This is no longer true with modern SSD devices and the optimization had no real benefit. Furthermore it caused increased fragmentation. The layout\nThe ssd_spread mount option attempts to allocate into bigger and aligned chunks of unused space, and may perform better on low-end SSDs. ssd_spread implies ssd, enabling all other SSD heuristics as well. The option nossd will disable all SSD options while nossd_spread only disables ssd_spread.",
	),
	docvalues.CreateEnumStringWithDoc(
		"treelog",
		"Enable the tree logging used for fsync and O_SYNC writes. The tree log stores changes without the need of a full filesystem sync. The log operations are flushed at sync and transaction commit. If the system crashes between two such syncs, the pending tree log operations are replayed during mount.\nWARNING:\nCurrently, the tree log is replayed even with a read-only mount! To disable that behaviour, also mount with nologreplay.\nThe tree log could contain new files/directories, these would not exist on a mounted filesystem if the log is not replayed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"notreelog",
		"Enable the tree logging used for fsync and O_SYNC writes. The tree log stores changes without the need of a full filesystem sync. The log operations are flushed at sync and transaction commit. If the system crashes between two such syncs, the pending tree log operations are replayed during mount.\nWARNING:\nCurrently, the tree log is replayed even with a read-only mount! To disable that behaviour, also mount with nologreplay.\nThe tree log could contain new files/directories, these would not exist on a mounted filesystem if the log is not replayed.",
	),
	docvalues.CreateEnumStringWithDoc(
		"usebackuproot",
		"Enable autorecovery attempts if a bad tree root is found at mount time. Currently this scans a backup list of several previous tree roots and tries to use the first readable. This can be used with read-only mounts as well.\nNOTE:\nThis option has replaced recovery.",
	),
	docvalues.CreateEnumStringWithDoc(
		"user_subvol_rm_allowed",
		"Allow subvolumes to be deleted by their respective owner. Otherwise, only the root user can do that.\nNOTE:\nHistorically, any user could create a snapshot even if he was not owner of the source subvolume, the subvolume deletion has been restricted for that reason. The subvolume creation has been restricted but this mount option is still required. This is a usability issue. Since 4.18, the [rmdir(2)](https://manpages.debian.org/testing/manpages-dev/rmdir.2.en.html) syscall can delete an empty subvolume j",
	),
	docvalues.CreateEnumStringWithDoc(
		"recovery",
		"This option has been replaced by usebackuproot and should not be used but will work on 4.5+ kernels.",
	),
	docvalues.CreateEnumStringWithDoc(
		"inode_cache",
		"The functionality has been removed in 5.11, any stale data created by previous use of the inode_cache option can be removed by btrfs check --clear-ino-cache.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noinode_cache",
		"The functionality has been removed in 5.11, any stale data created by previous use of the inode_cache option can be removed by btrfs check --clear-ino-cache.",
	),
}
