package commondocumentation

import docvalues "config-lsp/doc-values"

var DevptsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"This sets the owner or the group of newly created pseudo terminals to the specified values. When nothing is specified, they will be set to the UID and GID of the creating process. For example, if there is a tty group with GID 5, then gid=5 will cause newly created pseudo terminals to belong to the tty group.",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"This sets the owner or the group of newly created pseudo terminals to the specified values. When nothing is specified, they will be set to the UID and GID of the creating process. For example, if there is a tty group with GID 5, then gid=5 will cause newly created pseudo terminals to belong to the tty group.",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"mode",
		"Set the mode of newly created pseudo terminals to the specified value. The default is 0600. A value of mode=620 and gid=5 makes \"mesg y\" the default on newly created pseudo terminals.",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"ptmxmode",
		"Set the mode for the new ptmx device node in the devpts filesystem. With the support for multiple instances of devpts (see newinstance option above), each instance has a private ptmx node in the root of the devpts filesystem (typically /dev/pts/ptmx). For compatibility with older versions of the kernel, the default mode of the new ptmx node is 0000. ptmxmode=value specifies a more useful mode for the ptmx node and is highly recommended when the newinstance option is specified. This option is only implemented in Linux kernel versions starting with 2.6.29. Further, this option is valid only if CONFIG_DEVPTS_MULTIPLE_INSTANCES is enabled in the kernel configuration.",
	): docvalues.StringValue{},
}

var DevptsDocumentationEnums = []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"newinstance",
		"Create a private instance of the devpts filesystem, such that indices of pseudo terminals allocated in this new instance are independent of indices created in other instances of devpts. All mounts of devpts without this newinstance option share the same set of pseudo terminal indices (i.e., legacy mode). Each mount of devpts with the newinstance option has a private set of pseudo terminal indices. This option is mainly used to support containers in the Linux kernel. It is implemented in Linux kernel versions starting with 2.6.29. Further, this mount option is valid only if CONFIG_DEVPTS_MULTIPLE_INSTANCES is enabled in the kernel configuration. To use this option effectively, /dev/ptmx must be a symbolic link to pts/ptmx. See Documentation/filesystems/devpts.txt in the Linux kernel source tree for details.",
	),
}
