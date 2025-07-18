package formatting

import (
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var DefaultFormattingOptions = protocol.FormattingOptions{
	"tabSize":                float64(4),
	"insertSpaces":           false,
	"trimTrailingWhitespace": true,
}

// Following options available
// %s - a replacement for a string
// /!' - add a quote around the string if it contains spaces.
// /t - insert a tab / or spaces depending on the options
type FormatTemplate string

func (f FormatTemplate) Format(
	options protocol.FormattingOptions,
	a ...any,
) string {
	trimTrailingSpace := true

	if shouldTrim, found := options["trimTrailingWhitespace"]; found {
		trimTrailingSpace = shouldTrim.(bool)
	}

	value := ""
	value = fmt.Sprintf(
		string(
			f.replace("/t", getTab(options)),
		),
		a...,
	)

	if trimTrailingSpace {
		value = strings.TrimRight(value, " ")
		value = strings.TrimRight(value, "\t")
	}

	value = surroundWithQuotes(value)

	return value
}

func (f FormatTemplate) replace(format string, replacement string) FormatTemplate {
	value := string(f)
	currentIndex := 0

	for {
		position := strings.Index(value[currentIndex:], format)

		if position == -1 {
			break
		}

		position = position + currentIndex
		currentIndex = position

		if position == 0 || value[position-1] != '\\' {
			value = value[:position] + replacement + value[position+len(format):]
		}
	}

	return FormatTemplate(value)
}

func surroundWithQuotes(s string) string {
	value := s
	currentIndex := 0

	for {
		startPosition := strings.Index(value[currentIndex:], "/!'")

		if startPosition == -1 {
			break
		}

		startPosition = startPosition + currentIndex + 3
		currentIndex = startPosition

		endPosition := strings.Index(value[startPosition:], "/!'")

		if endPosition == -1 {
			break
		}

		endPosition = endPosition + startPosition
		currentIndex = endPosition

		innerValue := value[startPosition:endPosition]

		if innerValue[0] == '"' && innerValue[len(innerValue)-1] == '"' && (len(innerValue) >= 2 || innerValue[len(innerValue)-2] != '\\') {
			// Already surrounded
			value = value[:startPosition-3] + innerValue + value[endPosition+3:]
		} else if strings.Contains(innerValue, " ") {
			value = value[:startPosition-3] + "\"" + innerValue + "\"" + value[endPosition+3:]
		} else {
			value = value[:startPosition-3] + innerValue + value[endPosition+3:]
		}

		if endPosition+3 >= len(value) {
			break
		}
	}

	return value
}

func getTab(options protocol.FormattingOptions) string {
	tabSize := options["tabSize"].(float64)
	insertSpace := options["insertSpaces"].(bool)

	if insertSpace {
		return strings.Repeat(" ", int(tabSize))
	} else {
		return "\t"
	}
}
