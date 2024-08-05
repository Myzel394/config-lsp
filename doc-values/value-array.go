package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ArrayContainsDuplicatesError struct {
	Value string
}

func (e ArrayContainsDuplicatesError) Error() string {
	return fmt.Sprintf("'%s' is a duplicate value (and duplicates are not allowed)", e.Value)
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
	// If this function is nil, no duplicate check is done.
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

func (v ArrayValue) CheckIsValid(value string) []*InvalidValue {
	errors := []*InvalidValue{}
	values := strings.Split(value, v.Separator)

	if *v.DuplicatesExtractor != nil {
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

			duplicateValuesAsList := utils.FilterMapWhere(valuesOccurrences, func(_ string, value int) bool {
				return value > 1
			})
			duplicateValues := utils.KeysAsSet(duplicateValuesAsList)

			duplicateIndexStart := uint32(0)
			duplicateIndexEnd := uint32(0)

			currentIndex := uint32(0)
			for _, rawValue := range values {
				if _, found := duplicateValues[rawValue]; found {
					duplicateIndexStart = currentIndex
					duplicateIndexEnd = currentIndex + uint32(len(rawValue))

					errors = append(errors, &InvalidValue{
						Err: ArrayContainsDuplicatesError{
							Value: rawValue,
						},
						Start: duplicateIndexStart,
						End:   duplicateIndexEnd,
					})
				}
			}

			return errors
		}
	}

	currentIndex := uint32(0)
	for _, subValue := range values {
		newErrors := v.SubValue.CheckIsValid(subValue)

		if len(newErrors) > 0 {
			ShiftInvalidValues(currentIndex, newErrors)
		}

		errors = append(errors, newErrors...)

		currentIndex += uint32(len(subValue) + len(v.Separator))
	}

	return errors
}

func (v ArrayValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if cursor == 0 {
		return v.SubValue.FetchCompletions(line, cursor)
	}

	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor-1))

	if found {
		line = line[uint32(relativePosition):]
		cursor -= uint32(relativePosition)
	}

	return v.SubValue.FetchCompletions(line, cursor)
}
