package fstabdocumentation

import (
	docvalues "config-lsp/doc-values"
	"regexp"
)

var UuidField = docvalues.RegexValue{
	Regex: *regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`),
}
var LabelField = docvalues.RegexValue{
	Regex: *regexp.MustCompile(`\S+`),
}

var SpecField = docvalues.OrValue{
	Values: []docvalues.DeprecatedValue{
		// docvalues.PathValue{
		// 	RequiredType: docvalues.PathTypeFile & docvalues.PathTypeExistenceOptional,
		// },
		docvalues.KeyEnumAssignmentValue{
			Separator:       "=",
			ValueIsOptional: false,
			Values: map[docvalues.EnumString]docvalues.DeprecatedValue{
				docvalues.CreateEnumString("UUID"):      UuidField,
				docvalues.CreateEnumString("PARTUUID"):  UuidField,
				docvalues.CreateEnumString("LABEL"):     LabelField,
				docvalues.CreateEnumString("PARTLABEL"): LabelField,
			},
		},
	},
}
