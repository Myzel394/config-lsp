package fields

import (
	docvalues "config-lsp/doc-values"
)

var booleanEnumValue = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumString("yes"),
		docvalues.CreateEnumString("no"),
	},
}

var channelTimeoutExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")
var setEnvExtractor = docvalues.ExtractKeyDuplicatesExtractor("=")

func prefixPlusMinusCaret(values []docvalues.EnumString) docvalues.PrefixValue {
	return docvalues.PrefixValue{
		Prefixes: []docvalues.Prefix{
			{
				Prefix:  "+",
				Meaning: "Append to the default set",
			},
			{
				Prefix:  "-",
				Meaning: "Remove from the default set",
			},
			{
				Prefix:  "^",
				Meaning: "Place at the head of the default set",
			},
		},
		SubValue: docvalues.ArrayValue{
			Separator:           ",",
			DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
			RespectQuotes:       true,
			SubValue: docvalues.EnumValue{
				Values: values,
			},
		},
	}
}
