package fields

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
)

var BooleanField = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumString("yes"),
		docvalues.CreateEnumString("on"),
		docvalues.CreateEnumString("true"),
		docvalues.CreateEnumString("1"),

		docvalues.CreateEnumString("no"),
		docvalues.CreateEnumString("off"),
		docvalues.CreateEnumString("false"),
		docvalues.CreateEnumString("0"),
	},
}

var ColorField = docvalues.ArrayValue{
	DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
	Separator:           " ",
	SubValue: docvalues.OrValue{
		Values: []docvalues.DeprecatedValue{
			docvalues.NumberRangeValue(0, 255),
			docvalues.HexColorValue{Allow12Bit: true},
			docvalues.EnumValue{
				EnforceValues: true,
				Values: utils.MergeSlices(
					fetchColors(),
					fetchAttributes(),
					[]docvalues.EnumString{
						docvalues.CreateEnumString("normal"),
						docvalues.CreateEnumString("default"),
						docvalues.CreateEnumString("reset"),
					},
				),
			},
		},
	},
}

func fetchColors() []docvalues.EnumString {
	colors := []string{
		"black",
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
		"white",
	}

	enums := make([]docvalues.EnumString, 0)

	for _, color := range colors {
		enums = append(enums, docvalues.CreateEnumString(color))
		enums = append(enums, docvalues.CreateEnumString("bright"+color))
	}

	return enums
}

func fetchAttributes() []docvalues.EnumString {
	attributes := []string{
		"bold",
		"dim",
		"ul",
		"blink",
		"reverse",
		"italic",
		"strike",
	}

	enums := make([]docvalues.EnumString, 0)

	for _, attribute := range attributes {
		enums = append(enums, docvalues.CreateEnumString(attribute))
		enums = append(enums, docvalues.CreateEnumString("no"+attribute))
		enums = append(enums, docvalues.CreateEnumString("no-"+attribute))
	}

	return enums
}
