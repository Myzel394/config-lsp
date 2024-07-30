package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ArrayContainsDuplicatesError struct {
	Duplicates []string
}

func (e ArrayContainsDuplicatesError) Error() string {
	return fmt.Sprintf("The following values are duplicated: %s", strings.Join(e.Duplicates, ","))
}

var SimpleDuplicatesExtractor = func(value string) string {
	return value
}

var ExtractKeyDuplicatesExtractor = func(separator string) func(string) string {
	return func(value string) string {
		splitted := strings.Split(value, separator)

		if len(splitted) == 0 {
			return ""
		}

		return splitted[0]
	}
}

var DuplicatesAllowedExtractor func(string) string = nil

type ArrayValue struct {
	SubValue  Value
	Separator string
	// If this function is nil, no duplicate check is done
	// (value) => Extracted value
	// This is used to extract the value from the user input,
	// because you may want to preprocess the value before checking for duplicates
	DuplicatesExtractor *(func(string) string)
}

func (v ArrayValue) GetTypeDescription() []string {
	subValue := v.SubValue.(Value)

	return append(
		[]string{fmt.Sprintf("An Array separated by '%s' of:", v.Separator)},
		subValue.GetTypeDescription()...,
	)
}
func (v ArrayValue) CheckIsValid(value string) error {
	values := strings.Split(value, v.Separator)

	if v.DuplicatesExtractor != nil {
		valuesOccurrences := utils.SliceToMap(
			utils.Map(values, *v.DuplicatesExtractor),
			0,
		)

		// Only continue if there are actually duplicate values
		if len(values) != len(valuesOccurrences) {
			for _, duplicateRawValue := range values {
				duplicateValue := (*v.DuplicatesExtractor)(duplicateRawValue)
				valuesOccurrences[duplicateValue]++
			}

			duplicateValues := utils.FilterMapWhere(valuesOccurrences, func(_ string, value int) bool {
				return value > 1
			})

			return ArrayContainsDuplicatesError{
				Duplicates: utils.KeysOfMap(duplicateValues),
			}
		}
	}

	for _, subValue := range values {
		err := v.SubValue.CheckIsValid(subValue)

		if err != nil {
			return err
		}
	}

	return nil
}

func (v ArrayValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor-1))

	if found {
		line = line[uint32(relativePosition):]
		cursor -= uint32(relativePosition)
	}

	return v.SubValue.FetchCompletions(line, cursor)
}
