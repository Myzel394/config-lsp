package commondocumentation

import docvalues "config-lsp/doc-values"

var UsbfsDocumentationAssignable = docvalues.MergeKeyEnumAssignmentMaps(FatDocumentationAssignable, map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"devuid",
		"Set the owner of the device files in the usbfs filesystem (default: uid=gid=0)",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"devgid",
		"Set the group of the device files in the usbfs filesystem (default: uid=gid=0)",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"devmode",
		"Set the mode of the device files in the usbfs filesystem (default: 0644). The mode is given in octal.",
	): docvalues.StringValue{},

	docvalues.CreateEnumStringWithDoc(
		"busuid",
		"Set the owner of the bus files in the usbfs filesystem (default: uid=gid=0)",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"busgid",
		"Set the group of the bus files in the usbfs filesystem (default: uid=gid=0)",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"busmode",
		"Set the mode of the bus files in the usbfs filesystem (default: 0555). The mode is given in octal.",
	): docvalues.StringValue{},

	docvalues.CreateEnumStringWithDoc(
		"listuid",
		"Set the owner of the usbfs filesystem root directory (default: uid=gid=0)",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"listgid",
		"Set the group of the usbfs filesystem root directory (default: uid=gid=0)",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"listmode",
		"Set the mode of the usbfs filesystem root directory (default: 0555). The mode is given in octal.",
	): docvalues.StringValue{},
})

var UsbfsDocumentationEnums = []docvalues.EnumString{}
