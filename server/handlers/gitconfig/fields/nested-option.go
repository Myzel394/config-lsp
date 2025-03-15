package fields

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"regexp"
)

// In the gitconfig, there are certain options that can be nested.
// This is okay with the usual config-lsp setup, but some option
// allow custom names, such as: diff.<driver>.binary
// To handle those options correctly, we store nested options separately.
//
// This struct may only be used to define the options.
type NestedOptionDeclaration struct {
	// If true, the option is dynamic, like:
	// <driver>
	IsDynamic bool

	// If set, `SubOption` is the next option
	// If set, `Name` may not be set
	Options map[NormalizedOptionName]NestedOptionDeclaration
	// If set, `Name` is the documentation for the option
	// If set, `SubOption` may not be set
	Value *docvalues.DocumentationValue
}

func (n NestedOptionDeclaration) Parse(value string) (NestedOptionValue, error) {
	return NestedOptionValue{}, nil
}

// This struct may be used to parse nested options and store their values.
// How this works:
// We declare the options using `NestedOptionDeclaration`. They are read-only
// Then:
// 1. The antlr parser parses the text
// 2. The keys are manually parsed into a `NestedOptionValue`
// 3. Later, during the analyzer phase, it checks the validity of the keys, and
// 4. Adds the `Name` to the `nestedOptionValueValue`
type NestedOptionValue struct {
	Options []nestedOptionValueValue
}

type nestedOptionValueValue struct {
	common.LocationRange
	Option *NestedOptionDeclaration
	// Values can't be quoted, so we don't need a `ParsedString`
	Value string
}

var keyPartPattern = regexp.MustCompile(`[^.]+`)

func ParseNestedOptionToValue(
	text string,
	keyLocation common.LocationRange,
) NestedOptionValue {
	indexes := keyPartPattern.FindAllIndex([]byte(text), -1)
	options := make([]nestedOptionValueValue, 0)

	for _, index := range indexes {
		partText := text[index[0]:index[1]]
		partRange := keyLocation
		partRange.Start.Character = keyLocation.Start.Character + uint32(index[0])
		partRange.End.Character = keyLocation.Start.Character + uint32(index[1])

		options = append(
			options,
			nestedOptionValueValue{
				LocationRange: partRange,
				Value:         partText,
			},
		)
	}

	return NestedOptionValue{
		Options: options,
	}
}
