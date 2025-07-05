package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"regexp"
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
	SubValue  DeprecatedValue
	Separator string
	// If this function is nil, no duplicate check is done.
	// (value) => Extracted value
	// This is used to extract the value from the user input,
	// because you may want to preprocess the value before checking for duplicates
	DuplicatesExtractor *(func(string) string)

	// If true, array ArrayValue ignores the `Separator` if it's within quotes
	RespectQuotes bool
}

func (v ArrayValue) GetTypeDescription() []string {
	subValue := v.SubValue.(DeprecatedValue)

	return append(
		[]string{fmt.Sprintf("An Array separated by '%s' of:", v.Separator)},
		subValue.GetTypeDescription()...,
	)
}

// TODO: Add support for quotes
func (v ArrayValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	errors := []*InvalidValue{}
	var values []string

	if v.RespectQuotes {
		splitPattern := *regexp.MustCompile(fmt.Sprintf(`".+?"|[^%s]+`, v.Separator))

		values = splitPattern.FindAllString(value, -1)
	} else {
		values = strings.Split(value, v.Separator)
	}

	if *v.DuplicatesExtractor != nil {
		// Stores the values that are already found
		valueSet := map[string]struct{}{}

		currentIndex := uint32(0)

		for _, rawValue := range values {
			extractedValue := (*v.DuplicatesExtractor)(rawValue)
			valueLength := uint32(len(rawValue))

			if extractedValue == "" {
				// Skip empty values
				continue
			}

			if _, found := valueSet[extractedValue]; found {
				// This value is a duplicate, so we add an error
				errors = append(errors, &InvalidValue{
					Err: ArrayContainsDuplicatesError{
						Value: rawValue,
					},
					Start: currentIndex,
					End:   currentIndex + valueLength,
				})
			} else {
				valueSet[extractedValue] = struct{}{}
			}

			currentIndex += valueLength + uint32(len(v.Separator))
		}

		if len(errors) > 0 {
			// If there are errors, we return them immediately
			return errors
		}
	}

	currentIndex := uint32(0)
	for _, subValue := range values {
		newErrors := v.SubValue.DeprecatedCheckIsValid(subValue)

		if len(newErrors) > 0 {
			ShiftInvalidValues(currentIndex, newErrors)
		}

		errors = append(errors, newErrors...)

		currentIndex += uint32(len(subValue) + len(v.Separator))
	}

	return errors
}

func (v ArrayValue) getCurrentValue(line string, cursor uint32) (string, uint32) {
	if line == "" {
		return line, cursor
	}

	MIN := uint32(0)
	MAX := uint32(len(line) - 1)

	var cursorSearchStart = cursor
	var cursorSearchEnd = cursor

	var start uint32
	var end uint32

	// Hello,world,how,are,you
	// Hello,"world,how",are,you
	if v.RespectQuotes {
		quotes := utils.GetQuoteRanges(line)

		if len(quotes) > 0 {
			quote := quotes.GetQuoteForIndex(int(cursor))

			if quote != nil {
				cursorSearchStart = uint32(quote[0])
				cursorSearchEnd = uint32(quote[1])
			}
		}
	}

	// hello,w[o]rld,and,more
	// [h]ello,world
	// hello,[w]orld
	// hell[o],world
	// hello,worl[d]
	// hello,world[,]
	// hello[,]world,how,are,you
	relativePosition, found := utils.FindPreviousCharacter(
		line,
		v.Separator,
		int(cursorSearchStart),
	)

	if found {
		// + 1 to skip the separator
		start = min(
			MAX,
			uint32(relativePosition)+1,
		)
	} else {
		start = MIN
	}

	relativePosition, found = utils.FindNextCharacter(
		line,
		v.Separator,
		int(cursorSearchEnd),
	)

	if found {
		// - 1 to skip the separator
		end = max(
			MIN,
			uint32(relativePosition)-1,
		)
	} else {
		end = MAX
	}

	if cursor > end {
		// The user is typing a new (yet empty) value
		return "", 0
	}

	return line[start : end+1], cursor - start
}

func (v ArrayValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	value, cursor := v.getCurrentValue(line, cursor)

	return v.SubValue.DeprecatedFetchCompletions(value, cursor)
}

func (v ArrayValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	value, cursor := v.getCurrentValue(line, cursor)

	return v.SubValue.DeprecatedFetchHoverInfo(value, cursor)
}
