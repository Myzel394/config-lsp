package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ValueArrayContainsDuplicatesError struct {
	Value string
}

func (e ValueArrayContainsDuplicatesError) Error() string {
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

	// If true, the quotes will not be unwrapped and the raw string will be passed
	// to the subvalue
	PersistQuotes bool
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

	if v.DuplicatesExtractor != nil && *v.DuplicatesExtractor != nil {
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
					Err: ValueArrayContainsDuplicatesError{
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

func (v ArrayValue) getCurrentValue(line string, cursor common.CursorPosition) (string, common.CursorPosition) {
	index := max(uint32(cursor), 1) - 1

	if line == "" {
		return line, 0
	}

	MIN := uint32(0)
	MAX := uint32(len(line))

	var indexSearchStart = index
	var indexSearchEnd = index

	var start uint32
	var end uint32

	// Hello,world,how,are,you
	// Hello,"world,how",are,you
	if v.RespectQuotes {
		quotes := utils.GetQuoteRanges(line)

		if len(quotes) > 0 {
			quote := quotes.GetQuoteForIndex(int(index))

			if quote != nil {
				indexSearchStart = uint32(quote[0])
				indexSearchEnd = uint32(quote[1])
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
		int(indexSearchStart),
	)

	if found {
		// Edge case, two separators in a row
		// -> value is empty
		if relativePosition < (len(line)-1) && line[relativePosition+1] == v.Separator[0] {
			return "", 0
		}

		// Edge case, separator at start
		if relativePosition == 0 && cursor == 0 {
			return "", 0
		}

		// + 1 to skip the separator
		start = min(
			MAX,
			uint32(relativePosition)+1,
		)
		indexSearchEnd += 1
	} else {
		start = MIN
	}

	if indexSearchEnd >= uint32(len(line)) {
		end = MAX
	} else {

		relativePosition, found = utils.FindNextCharacter(
			line,
			v.Separator,
			int(indexSearchEnd),
		)

		if found {
			// - 1 to skip the separator
			end = max(
				MIN,
				uint32(relativePosition)-1,
			) + 1
		} else {
			end = MAX
		}

	}

	return line[start:end], cursor.ShiftHorizontal(-start)
}

func (v ArrayValue) unwrapQuotes(value string, cursor common.CursorPosition) (string, common.CursorPosition) {
	if v.RespectQuotes && !v.PersistQuotes && len(value) >= 1 {
		newValue := strings.TrimSuffix(strings.TrimPrefix(value, "\""), "\"")

		if value[0] == '"' && cursor <= 1 {
			return newValue, 0
		}

		if value[len(value)-1] == '"' && uint32(cursor) >= uint32(len(value))-1 {
			return newValue, common.CursorPosition(uint32(len(value)))
		}
	}

	// Nothing to do
	return value, cursor
}

func (v ArrayValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	relativeValue, relativeCursor := v.getCurrentValue(value, cursor)
	newValue, newCursor := v.unwrapQuotes(relativeValue, relativeCursor)

	return v.SubValue.FetchCompletions(newValue, newCursor)
}

func (v ArrayValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return nil
}
