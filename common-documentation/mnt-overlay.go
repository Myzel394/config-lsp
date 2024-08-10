package commondocumentation

import docvalues "config-lsp/doc-values"

var OverlayDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
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
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
			docvalues.CreateEnumString("follow"),
			docvalues.CreateEnumString("nofollow"),
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
		"When the underlying filesystems supports NFS export and the \"nfs_export\" feature is enabled, an overlay filesystem may be exported to NFS.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"xino",
		"The \"xino\" feature composes a unique object identifier from the real object st_ino and an underlying fsid index. The \"xino\" feature uses the high inode number bits for fsid, because the underlying filesystems rarely use the high inode number bits. In case the underlying inode number does overflow into the high xino bits, overlay filesystem will fall back to the non xino behavior for that inode.",
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
		"When metadata only copy up feature is enabled, overlayfs will only copy up metadata (as opposed to whole file), when a metadata specific operation like chown/chmod is performed. Full file will be copied up later when file is opened for WRITE operation.",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumString("on"),
			docvalues.CreateEnumString("off"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"volatile",
		"Volatile mounts are not guaranteed to survive a crash. It is strongly recommended that volatile mounts are only used if data written to the overlay can be recreated without significant effort.",
	): docvalues.StringValue{},
}

var OverlayDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"userxattr",
		"Use the \"user.overlay.\" xattr namespace instead of \"trusted.overlay.\". This is useful for unprivileged mounting of overlayfs.",
	),
	docvalues.CreateEnumStringWithDoc(
		"redirect_dir",
		"Redirects are enabled.",
	),
	docvalues.CreateEnumStringWithDoc(
		"index",
		"Inode index. If this feature is disabled and a file with multiple hard links is copied up, then this will \"break\" the link. Changes will not be propagated to other names referring to the same inode.",
	),
	docvalues.CreateEnumStringWithDoc(
		"uuid",
		"Can be used to replace UUID of the underlying filesystem in file handles with null, and effectively disable UUID checks. This can be useful in case the underlying disk is copied and the UUID of this copy is changed. This is only applicable if all lower/upper/work directories are on the same filesystem, otherwise it will fallback to normal behaviour.",
	),
	docvalues.CreateEnumStringWithDoc(
		"nfs_export",
		"When the underlying filesystems supports NFS export and the \"nfs_export\" feature is enabled, an overlay filesystem may be exported to NFS.",
	),
	docvalues.CreateEnumStringWithDoc(
		"xino",
		"The \"xino\" feature composes a unique object identifier from the real object st_ino and an underlying fsid index. The \"xino\" feature uses the high inode number bits for fsid, because the underlying filesystems rarely use the high inode number bits. In case the underlying inode number does overflow into the high xino bits, overlay filesystem will fall back to the non xino behavior for that inode.",
	),
	docvalues.CreateEnumStringWithDoc(
		"metacopy",
		"When metadata only copy up feature is enabled, overlayfs will only copy up metadata (as opposed to whole file), when a metadata specific operation like chown/chmod is performed. Full file will be copied up later when file is opened for WRITE operation.",
	),
	docvalues.CreateEnumStringWithDoc(
		"volatile",
		"Volatile mounts are not guaranteed to survive a crash. It is strongly recommended that volatile mounts are only used if data written to the overlay can be recreated without significant effort.",
	),
}
