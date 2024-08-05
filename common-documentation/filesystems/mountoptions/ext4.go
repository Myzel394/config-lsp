package commondocumentation

import docvalues "config-lsp/doc-values"

var zero = 0
var maxJournalIOPrio = 7

var Ext4DocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"inode_readahead_blks",
		"This tuning parameter controls the maximum number of inode table blocks that ext4's inode table readahead algorithm will pre-read into the buffer cache.  The value must be a power of 2. The default value is 32 blocks.",
	): docvalues.PowerOfTwoValue{},
	docvalues.CreateEnumStringWithDoc(
		"stripe",
		"Number of file system blocks that mballoc will try to use for allocation size and alignment. For RAID5/6 systems this should be the number of data disks * RAID chunk size in file system blocks.",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"max_batch_time",
		`Maximum amount of time ext4 should wait for additional file system operations to be batch together with a synchronous write operation. Since a synchronous write operation is going to force a commit and then a wait for the I/O complete, it doesn't cost much, and can be a huge throughput win, we wait for a small amount of time to see if any other transactions can piggyback on the synchronous write. The algorithm used is designed to automatically tune for the speed of the disk, by measuring the amount of time (on average) that it takes to finish committing a transaction. Call this time the "commit time".  If the time that the transaction has been running is less than the commit time, ext4 will try sleeping for the commit time to see if other operations will join the transaction. The commit time is capped by the max_batch_time, which defaults to 15000 µs (15 ms). This optimization can be turned off entirely by setting max_batch_time to 0.`,
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"min_batch_time",
		"This parameter sets the commit time (as described above) to be at least min_batch_time. It defaults to zero microseconds. Increasing this parameter may improve the throughput of multi-threaded, synchronous workloads on very fast disks, at the cost of increasing latency.",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"journal_ioprio",
		"The I/O priority (from 0 to 7, where 0 is the highest priority) which should be used for I/O operations submitted by kjournald2 during a commit operation.  This defaults to 3, which is a slightly higher priority than the default I/O priority.",
	): docvalues.NumberValue{Min: &zero, Max: &maxJournalIOPrio},
	docvalues.CreateEnumStringWithDoc(
		"init_itable",
		"The lazy itable init code will wait n times the number of milliseconds it took to zero out the previous block group's inode table. This minimizes the impact on system performance while the file system's inode table is being initialized.",
	): docvalues.PositiveNumberValue(),
	docvalues.CreateEnumStringWithDoc(
		"max_dir_size_kb",
		"This limits the size of the directories so that any attempt to expand them beyond the specified limit in kilobytes will cause an ENOSPC error. This is useful in memory-constrained environments, where a very large directory can cause severe performance problems or even provoke the Out Of Memory killer. (For example, if there is only 512 MB memory available, a 176 MB directory may seriously cramp the system's style.)",
	): docvalues.PositiveNumberValue(),
}

var Ext4DocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"journal_checksum",
		"The journal_checksum option enables checksumming of the journal transactions.  This will allow the recovery code in e2fsck and the kernel to detect corruption in the kernel. It is a compatible change and will be ignored by older kernels.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nojournal_checksum",
		"The journal_checksum option enables checksumming of the journal transactions.  This will allow the recovery code in e2fsck and the kernel to detect corruption in the kernel. It is a compatible change and will be ignored by older kernels.",
	),
	docvalues.CreateEnumStringWithDoc(
		"journal_async_commit",
		"Commit block can be written to disk without waiting for descriptor blocks. If enabled older kernels cannot mount the device.  This will enable 'journal_checksum' internally.",
	),
	docvalues.CreateEnumStringWithDoc(
		"barrier",
		`These mount options have the same effect as in ext3.  The mount options "barrier" and "nobarrier" are added for consistency with other ext4 mount options.
    The ext4 file system enables write barriers by default.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nobarrier",
		`These mount options have the same effect as in ext3.  The mount options "barrier" and "nobarrier" are added for consistency with other ext4 mount options.
    The ext4 file system enables write barriers by default.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"delalloc",
		"Deferring block allocation until write-out time.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodelalloc",
		"Disable delayed allocation. Blocks are allocated when data is copied from user to page cache.",
	),
	docvalues.CreateEnumStringWithDoc(
		"abort",
		"Simulate the effects of calling ext4_abort() for debugging purposes.  This is normally used while remounting a file system which is already mounted.",
	),
	docvalues.CreateEnumStringWithDoc(
		"auto_da_alloc",
		`Many broken applications don't use fsync() when replacing existing files via patterns such as

	fd = open("foo.new")/write(fd,...)/close(fd)/
	rename("foo.new", "foo")

	or worse yet

	fd = open("foo", O_TRUNC)/write(fd,...)/close(fd).

	If auto_da_alloc is enabled, ext4 will detect the replace- vi -rename and replace-via-truncate patterns and force that any delayed allocation blocks are allocated such that at the next journal commit, in the default data=ordered mode, the data blocks of the new file are forced to disk before the rename() operation is committed.  This provides roughly the same level of guarantees as ext3, and avoids the "zero-length" problem that can happen when a system crashes before the delayed allocation blocks are forced to disk.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noauto_da_alloc",
		`Many broken applications don't use fsync() when replacing existing files via patterns such as

	fd = open("foo.new")/write(fd,...)/close(fd)/
	rename("foo.new", "foo")

	or worse yet

	fd = open("foo", O_TRUNC)/write(fd,...)/close(fd).

	If auto_da_alloc is enabled, ext4 will detect the replace- vi -rename and replace-via-truncate patterns and force that any delayed allocation blocks are allocated such that at the next journal commit, in the default data=ordered mode, the data blocks of the new file are forced to disk before the rename() operation is committed.  This provides roughly the same level of guarantees as ext3, and avoids the "zero-length" problem that can happen when a system crashes before the delayed allocation blocks are forced to disk.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noinit_itable",
		"Do not initialize any uninitialized inode table blocks in the background. This feature may be used by installation CD's so that the install process can complete as quickly as possible; the inode table initialization process would then be deferred until the next time the file system is mounted.",
	),
	docvalues.CreateEnumStringWithDoc(
		"discard",
		"Controls whether ext4 should issue discard/TRIM commands to the underlying block device when blocks are freed. This is useful for SSD devices and sparse/thinly- provisioned LUNs, but it is off by default until sufficient testing has been done.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nodiscard",
		"Controls whether ext4 should issue discard/TRIM commands to the underlying block device when blocks are freed. This is useful for SSD devices and sparse/thinly- provisioned LUNs, but it is off by default until sufficient testing has been done.",
	),
	docvalues.CreateEnumStringWithDoc(
		"block_validity",
		"This option enables the in-kernel facility for tracking file system metadata blocks within internal data structures. This allows multi-block allocator and other routines to quickly locate extents which might overlap with file system metadata blocks. This option is intended for debugging purposes and since it negatively affects the performance, it is off by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"noblock_validity",
		"This option disables the in-kernel facility for tracking file system metadata blocks within internal data structures. This allows multi-block allocator and other routines to quickly locate extents which might overlap with file system metadata blocks. This option is intended for debugging purposes and since it negatively affects the performance, it is off by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"dioread_lock",
		"Controls whether or not ext4 should use the DIO read locking. If the dioread_nolock option is specified ext4 will allocate uninitialized extent before buffer write and convert the extent to initialized after IO completes. This approach allows ext4 code to avoid using inode mutex, which improves scalability on high speed storages. However this does not work with data journaling and dioread_nolock option will be ignored with kernel warning.  Note that dioread_nolock code path is only used for extent-based files.  Because of the restrictions this options comprises it is off by default (e.g. dioread_lock).",
	),
	docvalues.CreateEnumStringWithDoc(
		"dioread_nolock",
		"Controls whether or not ext4 should use the DIO read locking. If the dioread_nolock option is specified ext4 will allocate uninitialized extent before buffer write and convert the extent to initialized after IO completes. This approach allows ext4 code to avoid using inode mutex, which improves scalability on high speed storages. However this does not work with data journaling and dioread_nolock option will be ignored with kernel warning.  Note that dioread_nolock code path is only used for extent-based files.  Because of the restrictions this options comprises it is off by default (e.g. dioread_lock).",
	),
	docvalues.CreateEnumStringWithDoc(
		"i_version",
		"Enable 64-bit inode version support. This option is off by default.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nombcache",
		"This option disables use of mbcache for extended attribute deduplication. On systems where extended attributes are rarely or never shared between files, use of mbcache for deduplication adds unnecessary computational overhead.",
	),
	docvalues.CreateEnumStringWithDoc(
		"prjquota",
		"The prjquota mount option enables project quota support on the file system.  You need the quota utilities to actually enable and manage the quota system.  This mount option requires the project file system feature.",
	),
}
