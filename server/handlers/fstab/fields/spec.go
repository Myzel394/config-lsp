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

var sambaShareField = docvalues.RegexValue{
	// Match even a trailing slash and later show a diagnostic
	Regex: *regexp.MustCompile(`^//[^/]+?/.+$`),
}
var SpecField = docvalues.OrValue{
	Values: []docvalues.DeprecatedValue{
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
		sambaShareField,
		docvalues.PathValue{
			IsOptional:   false,
			RequiredType: docvalues.PathTypeFile,
		},
	},
}
