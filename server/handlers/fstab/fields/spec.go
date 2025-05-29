package fields

import (
	docvalues "config-lsp/doc-values"
	"regexp"
)

var UuidField = docvalues.RegexValue{
	// Can either be a UUID or UID
	Regex: *regexp.MustCompile(`(?i)([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}|[a-f0-9]{4}-[a-f0-9]{4})`),
}
var LabelField = docvalues.RegexValue{
	Regex: *regexp.MustCompile(`\S+`),
}

var SpecField = docvalues.OrValue{
	Values: []docvalues.DeprecatedValue{
		docvalues.PathValue{
			IsOptional: false,
			RequiredType: docvalues.PathTypeFile,
		},
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
