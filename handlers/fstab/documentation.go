package fstab

import (
	docvalues "config-lsp/doc-values"
	"regexp"
)

var uuidField = docvalues.RegexValue{
	Regex: *regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`),
}
var labelField = docvalues.RegexValue{
	Regex: *regexp.MustCompile(`\S+`),
}

var specField = docvalues.OrValue{
	Values: []docvalues.Value{
		// docvalues.PathValue{
		// 	RequiredType: docvalues.PathTypeFile & docvalues.PathTypeExistenceOptional,
		// },
		docvalues.KeyValueAssignmentValue{
			Separator:       "=",
			ValueIsOptional: false,
			Key: docvalues.EnumValue{
				EnforceValues: true,
				Values: []docvalues.EnumString{
					docvalues.CreateEnumString("UUID"),
					docvalues.CreateEnumString("LABEL"),
					docvalues.CreateEnumString("PARTUUID"),
					docvalues.CreateEnumString("PARTLABEL"),
				},
			},
			Value: docvalues.CustomValue{
				FetchValue: func(rawContext docvalues.CustomValueContext) docvalues.Value {
					context := rawContext.(docvalues.KeyValueAssignmentContext)

					switch context.SelectedKey {
					case "UUID":
					case "PARTUUID":
						return uuidField
					case "LABEL":
					case "PARTLABEL":
						return labelField
					}

					return docvalues.StringValue{}
				},
			},
		},
	},
}

var mountPointField = docvalues.OrValue{
	Values: []docvalues.Value{
		docvalues.EnumValue{
			Values: []docvalues.EnumString{
				{
					InsertText:      "none",
					DescriptionText: "none",
					Documentation:   "Specify that the filesystem should be treated as swap space",
				},
			},
		},
		docvalues.RegexValue{
			Regex: *regexp.MustCompile(`\S+`),
		},
	},
}
