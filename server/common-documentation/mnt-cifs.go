package commondocumentation

import docvalues "config-lsp/doc-values"

var CifsDocumentationAssignable = map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"user",
		`specifies the username to connect as. If this is not given, then the environment variable USER is used. This option can also take the form "user%password" or "workgroup/user" or "workgroup/user%password" to allow the password and workgroup to be specified as part of the username.
Note
The cifs vfs accepts the parameter user=, or for users familiar with smbfs it accepts the longer form of the parameter username=. Similarly the longer smbfs style parameter names may be accepted as synonyms for the shorter cifs parameters pass=,dom= and cred=.`,
	): docvalues.UserValue("", false),
	docvalues.CreateEnumStringWithDoc(
		"password",
		`specifies the CIFS password. If this option is not given then the environment variable PASSWD is used. If the password is not specified directly or indirectly via an argument to mount, mount.cifs will prompt for a password, unless the guest option is specified.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"credentials",
		`specifies a file that contains a username and/or password and optionally the name of the workgroup. The format of the file is:
username=value

password=value

domain=value

This is preferred over having passwords in plaintext in a shared file, such as /etc/fstab. Be sure to protect any credentials file properly.`,
	): docvalues.PathValue{
		RequiredType: docvalues.PathTypeFile,
	},
	docvalues.CreateEnumStringWithDoc(
		"uid",
		`sets the uid that will own all files or directories on the mounted filesystem when the server does not provide ownership information. It may be specified as either a username or a numeric uid. When not specified, the default is uid 0. The mount.cifs helper must be at version 1.10 or higher to support specifying the uid in non-numeric form. See the section on FILE AND DIRECTORY OWNERSHIP AND PERMISSIONS below for more information.`,
	): docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.UIDValue{
				EnforceUsingExisting: false,
			},
			docvalues.UserValue("", false),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"cruid",
		`sets the uid of the owner of the credentials cache. This is primarily useful with sec=krb5. The default is the real uid of the process performing the mount. Setting this parameter directs the upcall to look for a credentials cache owned by that user.`,
	): docvalues.UIDValue{
		EnforceUsingExisting: false,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		`sets the gid that will own all files or directories on the mounted filesystem when the server does not provide ownership information. It may be specified as either a groupname or a numeric gid. When not specified, the default is gid 0. The mount.cifs helper must be at version 1.10 or higher to support specifying the gid in non-numeric form. See the section on FILE AND DIRECTORY OWNERSHIP AND PERMISSIONS below for more information.`,
	): docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.GIDValue{
				EnforceUsingExisting: false,
			},
			docvalues.GroupValue("", false),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"port",
		`sets the port number on which the client will attempt to contact the CIFS server. If this value is specified, look for an existing connection with this port, and use that if one exists. If one doesn't exist, try to create a new connection on that port. If that connection fails, return an error. If this value isn't specified, look for an existing connection on port 445 or 139. If no such connection exists, try to connect on port 445 first and then port 139 if that fails. Return an error if both fail.`,
	): docvalues.NumberRangeValue(1, 65535),
	docvalues.CreateEnumStringWithDoc(
		"servernetbiosname",
		`Specify the server netbios name (RFC1001 name) to use when attempting to setup a session to the server. Although rarely needed for mounting to newer servers, this option is needed for mounting to some older servers (such as OS/2 or Windows 98 and Windows ME) since when connecting over port 139 they, unlike most newer servers, do not support a default server name. A server name can be up to 15 characters long and is usually uppercased.`,
	// TODO: Add regex check for RFC1001 name
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"servern",
		"Synonym for servernetbiosname.",
	): docvalues.StringValue{},
	// TODO: Show warning when port is not 193
	docvalues.CreateEnumStringWithDoc(
		"netbiosname",
		`When mounting to servers via port 139, specifies the RFC1001 source name to use to represent the client netbios machine name when doing the RFC1001 netbios session initialize.`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"file_mode",
		`If the server does not support the CIFS Unix extensions this overrides the default file mode.`,
	): docvalues.MaskModeValue{},
	docvalues.CreateEnumStringWithDoc(
		"dir_mode",
		`If the server does not support the CIFS Unix extensions this overrides the default mode for directories.`,
	): docvalues.MaskModeValue{},
	docvalues.CreateEnumStringWithDoc(
		"ip",
		`sets the destination IP address. This option is set automatically if the server name portion of the requested UNC name can be resolved so rarely needs to be specified by the user.`,
	): docvalues.IPAddressValue{
		AllowIPv4: true,
		AllowIPv6: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"domain",
		`sets the domain (workgroup) of the user`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"cache",
		`Cache mode. See the section below on CACHE COHERENCY for details.
The default in kernels prior to 3.7 was "loose". As of kernel 3.7 the default is "strict".`,
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("none", "do not cache file data at all"),
			docvalues.CreateEnumStringWithDoc("strict", "follow the CIFS/SMB2 protocol strictly"),
			docvalues.CreateEnumStringWithDoc("loose", "allow loose caching semantics"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"backupuid",
		`Restrict access to files with the backup intent to a user. Either a name or an id must be provided as an argument, there are no default values.
See section ACCESSING FILES WITH BACKUP INTENT for more details`,
	): docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.UIDValue{
				EnforceUsingExisting: false,
			},
			docvalues.UserValue("", false),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"backupgid",
		`Restrict access to files with the backup intent to a group. Either a name or an id must be provided as an argument, there are no default values.
See section ACCESSING FILES WITH BACKUP INTENT for more details`,
	): docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.GIDValue{
				EnforceUsingExisting: false,
			},
			docvalues.GroupValue("", false),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"sec",
		`Security mode.
If the server requires signing during protocol negotiation, then it may be enabled automatically. Packet signing may also be enabled automatically if it's enabled in /proc/fs/cifs/SecurityFlags.`,
	): docvalues.EnumValue{
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc("none", "attempt to connection as a null user (no name)"),
			docvalues.CreateEnumStringWithDoc("krb5", "Use Kerberos version 5 authentication"),
			docvalues.CreateEnumStringWithDoc("krb5i", "Use Kerberos authentication and forcibly enable packet signing"),
			docvalues.CreateEnumStringWithDoc("ntlm", "Use NTLM password hashing (default)"),
			docvalues.CreateEnumStringWithDoc("ntlmi", "Use NTLM password hashing and force packet signing"),
			docvalues.CreateEnumStringWithDoc("ntlmv2", "Use NTLMv2 password hashing"),
			docvalues.CreateEnumStringWithDoc("ntlmv2i", "Use NTLMv2 password hashing and force packet signing"),
			docvalues.CreateEnumStringWithDoc("ntlmssp", "Use NTLMv2 password hashing encapsulated in Raw NTLMSSP message"),
			docvalues.CreateEnumStringWithDoc("ntlmsspi", "Use NTLMv2 password hashing encapsulated in Raw NTLMSSP message, and force packet signing"),
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"rsize",
		`default network read size (usually 16K). The client currently can not use rsize larger than CIFSMaxBufSize. CIFSMaxBufSize defaults to 16K and may be changed (from 8K to the maximum kmalloc size allowed by your kernel) at module install time for cifs.ko. Setting CIFSMaxBufSize to a very large value will cause cifs to use more memory and may reduce performance in some cases. To use rsize greater than 127K (the original cifs protocol maximum) also requires that the server support a new Unix Capability flag (for very large read) which some newer servers (e.g. Samba 3.0.26 or later) do. rsize can be set from a minimum of 2048 to a maximum of 130048 (127K or CIFSMaxBufSize, whichever is smaller)`,
	): docvalues.DataAmountValue{
		AllowedUnits: map[rune]struct{}{
			'k': {},
			'm': {},
		},
		AllowByteSuffix: false,
		Base:            docvalues.DataAmountValueBase1024,
		AllowDecimal:    false,
		Validator:       docvalues.CreateDARangeValidator("2048", "127K", docvalues.DataAmountValueBase1024),
	},
	docvalues.CreateEnumStringWithDoc(
		"wsize",
		`Maximum amount of data that the kernel will send in a write request in bytes. Prior to RHEL6.2 kernels, the default and maximum was 57344 (14 * 4096 pages). As of RHEL6.2, the default depends on whether the client and server negotiate large writes via POSIX extensions. If they do then the default is 1M, and the maximum allowed is 16M. If they do not, then the default is 65536 and the maximum allowed is 131007.
Note that this value is just a starting point for negotiation. The client and server may negotiate this size downward according to the server's capabilities.`,
	): docvalues.NumberRangeValue(1, 131007),
	docvalues.CreateEnumStringWithDoc(
		"actimeo",
		`The time (in seconds) that the CIFS client caches attributes of a file or directory before it requests attribute information from a server. During this period the changes that occur on the server remain undetected until the client checks the server again.
By default, the attribute cache timeout is set to 1 second. This means more frequent on-the-wire calls to the server to check whether attributes have changed which could impact performance. With this option users can make a tradeoff between performance and cache metadata correctness, depending on workload needs. Shorter timeouts mean better cache coherency, but frequent increased number of calls to the server. Longer timeouts mean a reduced number of calls to the server but looser cache coherency. The actimeo value is a positive integer that can hold values between 0 and a maximum value of 2^30 * HZ (frequency of timer interrupt) setting.`,
	): docvalues.NumberRangeValue(0, 1073741823),
	docvalues.CreateEnumStringWithDoc(
		"prefixpath",
		`It's possible to mount a subdirectory of a share. The preferred way to do this is to append the path to the UNC when mounting. However, it's also possible to do the same by setting this option and providing the path there.`,
	): docvalues.PathValue{
		RequiredType: docvalues.PathTypeDirectory,
		IsOptional:   true,
	},

	// This option is listed as having no arg, but I'm pretty sure it does
	docvalues.CreateEnumStringWithDoc(
		"iocharset",
		`Charset used to convert local path names to and from Unicode. Unicode is used by default for network path names if the server supports it. If iocharset is not specified then the nls_default specified during the local client kernel build will be used. If server does not support Unicode, this parameter is unused.`,
	): docvalues.EnumValue{
		EnforceValues: true,
		Values:        AvailableCharsets,
	},
}

var CifsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"forceuid",
		`instructs the client to ignore any uid provided by the server for files and directories and to always assign the owner to be the value of the uid= option. See the section on FILE AND DIRECTORY OWNERSHIP AND PERMISSIONS below for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"forcegid",
		`instructs the client to ignore any gid provided by the server for files and directories and to always assign the owner to be the value of the gid= option. See the section on FILE AND DIRECTORY OWNERSHIP AND PERMISSIONS below for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"guest",
		`don't prompt for a password`,
	),
	docvalues.CreateEnumStringWithDoc(
		"ro",
		`mount read-only`,
	),
	docvalues.CreateEnumStringWithDoc(
		"rw",
		`mount read-write`,
	),
	docvalues.CreateEnumStringWithDoc(
		"setuids",
		`If the CIFS Unix extensions are negotiated with the server the client will attempt to set the effective uid and gid of the local process on newly created files, directories, and devices (create, mkdir, mknod). If the CIFS Unix Extensions are not negotiated, for newly created files and directories instead of using the default uid and gid specified on the the mount, cache the new file's uid and gid locally which means that the uid for the file can change when the inode is reloaded (or the user remounts the share).`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nosetuids",
		`The client will not attempt to set the uid and gid on on newly created files, directories, and devices (create, mkdir, mknod) which will result in the server setting the uid and gid to the default (usually the server uid of the user who mounted the share). Letting the server (rather than the client) set the uid and gid is the default.If the CIFS Unix Extensions are not negotiated then the uid and gid for new files will appear to be the uid (gid) of the mounter or the uid (gid) parameter specified on the mount.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"perm",
		`Client does permission checks (vfs_permission check of uid and gid of the file against the mode and desired operation), Note that this is in addition to the normal ACL check on the target machine done by the server software. Client permission checking is enabled by default.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noperm",
		`Client does not do permission checks. This can expose files on this mount to access by other users on the local client system. It is typically only needed when the server supports the CIFS Unix Extensions but the UIDs/GIDs on the client and server system do not match closely enough to allow access by the user doing the mount. Note that this does not affect the normal ACL check on the target machine done by the server software (of the server ACL against the user name provided at mount time).`,
	),
	docvalues.CreateEnumStringWithDoc(
		"dynperm",
		`Instructs the server to maintain ownership and permissions in memory that can't be stored on the server. This information can disappear at any time (whenever the inode is flushed from the cache), so while this may help make some applications work, it's behavior is somewhat unreliable. See the section below on FILE AND DIRECTORY OWNERSHIP AND PERMISSIONS for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"directio",
		`Do not do inode data caching on files opened on this mount. This precludes mmaping files on this mount. In some cases with fast networks and little or no caching benefits on the client (e.g. when the application is doing large sequential reads bigger than page size without rereading the same data) this can provide better performance than the default behavior which caches reads (readahead) and writes (writebehind) through the local Linux client pagecache if oplock (caching token) is granted and held.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"strictcache",
		`Use for switching on strict cache mode. In this mode the client reads from the cache all the time it has Oplock Level II, otherwise - read from the server. As for write - the client stores a data in the cache in Exclusive Oplock case, otherwise - write directly to the server.
This option is will be deprecated in 3.7. Users should use cache=strict instead on more recent kernels.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"rwpidforward",
		`Forward pid of a process who opened a file to any read or write operation on that file. This prevent applications like WINE from failing on read and write if we use mandatory brlock style.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"mapchars",
		`Translate six of the seven reserved characters (not backslash, but including the colon, question mark, pipe, asterik, greater than and less than characters) to the remap range (above 0xF000), which also allows the CIFS client to recognize files created with such characters by Windows's POSIX emulation. This can also be useful when mounting to most versions of Samba (which also forbids creating and opening files whose names contain any of these seven characters). This has no effect if the server does not support Unicode on the wire. Please note that the files created with mapchars mount option may not be accessible if the share is mounted without that option.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nomapchars",
		`Do not translate any of these seven characters (default)`,
	),
	// TODO: Show a warning when this is used
	docvalues.CreateEnumStringWithDoc(
		"intr",
		`currently unimplemented`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nointr",
		`(default) currently unimplemented`,
	),
	docvalues.CreateEnumStringWithDoc(
		"hard",
		`The program accessing a file on the cifs mounted file system will hang when the server crashes.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"soft",
		`(default) The program accessing a file on the cifs mounted file system will not hang when the server crashes and will return errors to the user application.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noacl",
		`Do not allow POSIX ACL operations even if server would support them.
The CIFS client can get and set POSIX ACLs (getfacl, setfacl) to Samba servers version 3.0.10 and later. Setting POSIX ACLs requires enabling both CIFS_XATTR and then CIFS_POSIX support in the CIFS configuration options when building the cifs module. POSIX ACL support can be disabled on a per mount basis by specifying "noacl" on mount.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"cifsacl",
		`This option is used to map CIFS/NTFS ACLs to/from Linux permission bits, map SIDs to/from UIDs and GIDs, and get and set Security Descriptors.
See sections on CIFS/NTFS ACL, SID/UID/GID MAPPING, SECURITY DESCRIPTORS for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nocase",
		`Request case insensitive path name matching (case sensitive is the default if the server suports it).`,
	),
	// TODO: Show grayed text when this is used
	docvalues.CreateEnumStringWithDoc(
		"ignorecase",
		`Synonym for nocase.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nobrl",
		`Do not send byte range lock requests to the server. This is necessary for certain applications that break with cifs style mandatory byte range locks (and most cifs servers do not yet support requesting advisory byte range locks).`,
	),
	docvalues.CreateEnumStringWithDoc(
		"sfu",
		`When the CIFS Unix Extensions are not negotiated, attempt to create device files and fifos in a format compatible with Services for Unix (SFU). In addition retrieve bits 10-12 of the mode via the SETFILEBITS extended attribute (as SFU does). In the future the bottom 9 bits of the mode mode also will be emulated using queries of the security descriptor (ACL). [NB: requires version 1.39 or later of the CIFS VFS. To recognize symlinks and be able to create symlinks in an SFU interoperable form requires version 1.40 or later of the CIFS VFS kernel module.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"serverino",
		`Use inode numbers (unique persistent file identifiers) returned by the server instead of automatically generating temporary inode numbers on the client. Although server inode numbers make it easier to spot hardlinked files (as they will have the same inode numbers) and inode numbers may be persistent (which is userful for some sofware), the server does not guarantee that the inode numbers are unique if multiple server side mounts are exported under a single share (since inode numbers on the servers might not be unique if multiple filesystems are mounted under the same shared higher level directory). Note that not all servers support returning server inode numbers, although those that support the CIFS Unix Extensions, and Windows 2000 and later servers typically do support this (although not necessarily on every local server filesystem). Parameter has no effect if the server lacks support for returning inode numbers or equivalent. This behavior is enabled by default.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noserverino",
		`Client generates inode numbers itself rather than using the actual ones from the server.
See section INODE NUMBERS for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nounix",
		`Disable the CIFS Unix Extensions for this mount. This can be useful in order to turn off multiple settings at once. This includes POSIX acls, POSIX locks, POSIX paths, symlink support and retrieving uids/gids/mode from the server. This can also be useful to work around a bug in a server that supports Unix Extensions.
See section INODE NUMBERS for more information.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"nouser_xattr",
		`(default) Do not allow getfattr/setfattr to get/set xattrs, even if server would support it otherwise.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"fsc",
		`Enable local disk caching using FS-Cache for CIFS. This option could be useful to improve performance on a slow link, heavily loaded server and/or network where reading from the disk is faster than reading from the server (over the network). This could also impact the scalability positively as the number of calls to the server are reduced. But, be warned that local caching is not suitable for all workloads, for e.g., read-once type workloads. So, you need to consider carefully the situation/workload before using this option. Currently, local disk caching is enabled for CIFS files opened as read-only.
NOTE: This feature is available only in the recent kernels that have been built with the kernel config option CONFIG_CIFS_FSCACHE. You also need to have cachefilesd daemon installed and running to make the cache operational.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"multiuser",
		`Map user accesses to individual credentials when accessing the server. By default, CIFS mounts only use a single set of user credentials (the mount credentials) when accessing a share. With this option, the client instead creates a new session with the server using the user's credentials whenever a new user accesses the mount. Further accesses by that user will also use those credentials. Because the kernel cannot prompt for passwords, multiuser mounts are limited to mounts using sec= options that don't require passwords.
With this change, it's feasible for the server to handle permissions enforcement, so this option also implies "noperm". Furthermore, when unix extensions aren't in use and the administrator has not overriden ownership using the uid= or gid= options, ownership of files is presented as the current user accessing the share.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"noposixpaths",
		`If unix extensions are enabled on a share, then the client will typically allow filenames to include any character besides '/' in a pathname component, and will use forward slashes as a pathname delimiter. This option prevents the client from attempting to negotiate the use of posix-style pathnames to the server.`,
	),
	docvalues.CreateEnumStringWithDoc(
		"posixpaths",
		`Inverse of noposixpaths.`,
	),
}
