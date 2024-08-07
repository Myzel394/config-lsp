package commondocumentation

import docvalues "config-lsp/doc-values"

var AdfsDocumentationAssignable = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"uid",
		"Set the owner of the files in the filesystem (default: uid=0).",
	): docvalues.UIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"gid",
		"Set the group of the files in the filesystem (default: gid=0).",
	): docvalues.GIDValue{
		EnforceUsingExisting: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"ownmask",
		"Set the permission mask for ADFS 'owner' permissions (default: 0700).",
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"othmask",
		"Set the permission mask for ADFS 'other' permissions (default: 0077).",
	): docvalues.StringValue{},
}

var AdfsDocumentationEnums = []docvalues.EnumString{}
