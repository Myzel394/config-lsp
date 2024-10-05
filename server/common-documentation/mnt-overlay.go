package commondocumentation

import docvalues "config-lsp/doc-values"

var OverlayDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"lowerdir",
		"Any filesystem, does not need to be on a writable filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"upperdir",
		"The upperdir is normally on a writable filesystem.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"workdir",
		"The workdir needs to be an empty directory on the same filesystem as upperdir.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"redirect_dir",
		"If the redirect_dir feature is enabled, then the directory will be copied up (but not the contents). Then the \"{trusted|user}.overlay.redirect\" extended attribute is set to the path of the original location from the root of the overlay. Finally the directory is moved to the new location.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"on",
				"Redirects are enabled.",
			),
			docvalues.CreateEnumStringWithDoc(
				"off",
				"Redirects are not created and only followed if \"redirect_always_follow\" feature is enabled in the kernel/module config.",
			),
			docvalues.CreateEnumStringWithDoc(
				"follow",
				"Redirects are not created, but followed.",
			),
			docvalues.CreateEnumStringWithDoc(
				"nofollow",
				"Redirects are not created and not followed (equivalent to \"redirect_dir=off\" if \"redirect_always_follow\" feature is not enabled).",
			),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"index",
		"Inode index. If this feature is disabled and a file with multiple hard links is copied up, then this will \"break\" the link. Changes will not be propagated to other names referring to the same inode.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"uuid",
		"Can be used to replace UUID of the underlying filesystem in file handles with null, and effectively disable UUID checks. This can be useful in case the underlying disk is copied and the UUID of this copy is changed. This is only applicable if all lower/upper/work directories are on the same filesystem, otherwise it will fallback to normal behaviour.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"nfs_export",
		`When encoding a file handle from an overlay filesystem object, the following rules apply

   •   For a non-upper object, encode a lower file handle from lower inode

   •   For an indexed object, encode a lower file handle from copy_up origin

   •   For a pure-upper object and for an existing non-indexed upper object, encode an upper file handle from upper inode

The encoded overlay file handle includes

   •   Header including path type information (e.g. lower/upper)

   •   UUID of the underlying filesystem

   •   Underlying filesystem encoding of underlying inode

This encoding format is identical to the encoding format of file handles that are stored in extended attribute "{trusted|user}.overlay.origin". When decoding an overlay file handle, the following steps are followed

   •   Find underlying layer by UUID and path type information.

   •   Decode the underlying filesystem file handle to underlying dentry.

   •   For a lower file handle, lookup the handle in index directory by name.

   •   If a whiteout is found in index, return ESTALE. This represents an overlay object that was deleted after its file handle was encoded.

   •   For a non-directory, instantiate a disconnected overlay dentry from the decoded underlying dentry, the path type and index inode, if found.

   •   For a directory, use the connected underlying decoded dentry, path type and index, to lookup a connected overlay dentry.

Decoding a non-directory file handle may return a disconnected dentry. copy_up of that disconnected dentry will create an upper index entry with no upper alias.

When overlay filesystem has multiple lower layers, a middle layer directory may have a "redirect" to lower directory. Because middle layer "redirects" are not indexed, a lower file handle that was encoded from the "redirect" origin directory, cannot be used to find the middle or upper layer directory. Similarly, a lower file handle that was encoded from a descendant of the "redirect" origin directory, cannot be used to reconstruct a connected overlay path. To mitigate the cases of directories that cannot be decoded from a lower file handle, these directories are copied up on encode and encoded as an upper file handle. On an overlay filesystem with no upper layer this mitigation cannot be used NFS export in this setup requires turning off redirect follow (e.g.
"redirect_dir=nofollow").

The overlay filesystem does not support non-directory connectable file handles, so exporting with the subtree_check exportfs configuration will cause failures to lookup files over NFS.

When the NFS export feature is enabled, all directory index entries are verified on mount time to check that upper file handles are not stale. This verification may cause significant overhead in some cases.

Note: the mount options index=off,nfs_export=on are conflicting for a read-write mount and will result in an error.`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"xino",
		`The "xino" feature composes a unique object identifier from the real object st_ino and an underlying fsid index. The "xino" feature uses the high inode number bits for fsid, because the underlying filesystems rarely use the high inode number bits. In case the underlying inode number does overflow into the high xino bits, overlay filesystem will fall back to the non xino behavior for that inode.

    For a detailed description of the effect of this option please refer to https://docs.kernel.org/filesystems/overlayfs.html`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
			docvalues.CreateEnumString("auto"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"metacopy",
		`When metadata only copy up feature is enabled, overlayfs will only copy up metadata (as opposed to whole file), when a metadata specific operation like chown/chmod is performed. Full file will be copied up later when file is opened for WRITE operation.

    In other words, this is delayed data copy up operation and data is copied up when there is a need to actually modify data.`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
}

var OverlayDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"userxattr",
		"Use the \"user.overlay.\" xattr namespace instead of \"trusted.overlay.\". This is useful for unprivileged mounting of overlayfs.",
	),
	docvalues.CreateEnumStringWithDoc(
		"volatile",
		`Volatile mounts are not guaranteed to survive a crash. It is strongly recommended that volatile mounts are only used if data written to the overlay can be recreated without significant effort.

    The advantage of mounting with the "volatile" option is that all forms of sync calls to the upper filesystem are omitted.

    In order to avoid a giving a false sense of safety, the syncfs (and fsync) semantics of volatile mounts are slightly different than that of the rest of VFS. If any writeback error occurs on the upperdir’s filesystem after a volatile mount takes place, all sync functions will return an error. Once this condition is reached, the filesystem will not recover, and every subsequent sync call will return an error, even if the upperdir has not experience a new error since the last sync call.

    When overlay is mounted with "volatile" option, the directory "$workdir/work/incompat/volatile" is created. During next mount, overlay checks for this directory and refuses to mount if present. This is a strong indicator that user should throw away upper and work directories and create fresh one. In very limited cases where the user knows that the system has not crashed and contents of upperdir are intact, The "volatile" directory can be removed.
`,
	),
}
