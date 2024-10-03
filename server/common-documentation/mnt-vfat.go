package commondocumentation

import docvalues "config-lsp/doc-values"

var VfatDocumentationAssignable = docvalues.MergeKeyEnumAssignmentMaps(FatDocumentationAssignable, map[docvalues.EnumString]docvalues.DeprecatedValue{
	docvalues.CreateEnumStringWithDoc(
		"shortname",
		"Defines the behavior for creation and display of filenames which fit into 8.3 characters. If a long name for a file exists, it will always be the preferred one for display",
	): docvalues.EnumValue{
		EnforceValues: true,
		Values: []docvalues.EnumString{
			docvalues.CreateEnumStringWithDoc(
				"lower",
				"Force the short name to lower case upon display; store a long name when the short name is not all upper case.",
			),
			docvalues.CreateEnumStringWithDoc(
				"win95",
				"Force the short name to upper case upon display; store a long name when the short name is not all upper case.",
			),
			docvalues.CreateEnumStringWithDoc(
				"winnt",
				"Display the short name as is; store a long name when the short name is not all lower case or all upper case.",
			),
			docvalues.CreateEnumStringWithDoc(
				"mixed",
				"Display the short name as is; store a long name when the short name is not all upper case. This mode is the default since Linux 2.6.32.",
			),
		},
	},
})

var VfatDocumentationEnums = append(FatDocumentationEnums, []docvalues.EnumString{
	docvalues.CreateEnumStringWithDoc(
		"uni_xlate",
		"Translate unhandled Unicode characters to special escaped sequences. This lets you backup and restore filenames that are created with any Unicode characters. Without this option, a '?' is used when no translation is possible. The escape character is ':' because it is otherwise invalid on the vfat filesystem. The escape sequence that gets used, where u is the Unicode character, is: ':', (u & 0x3f), ((u>>6) & 0x3f), (u>>12).",
	),
	docvalues.CreateEnumStringWithDoc(
		"posix",
		"Allow two files with names that only differ in case. This option is obsolete.",
	),
	docvalues.CreateEnumStringWithDoc(
		"monumtail",
		"First try to make a short name without sequence number, before trying name~num.ext.",
	),
	docvalues.CreateEnumStringWithDoc(
		"utf8",
		"UTF8 is the filesystem safe 8-bit encoding of Unicode that is used by the console. It can be enabled for the filesystem with this option or disabled with utf8=0, utf8=no or utf8=false. If uni_xlate gets set, UTF8 gets disabled.",
	),
}...)
