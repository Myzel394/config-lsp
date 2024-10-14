package fields

import (
	docvalues "config-lsp/doc-values"
	"regexp"
)

var MountPointField = docvalues.OrValue{
	Values: []docvalues.DeprecatedValue{
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
