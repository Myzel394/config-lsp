package formatting

import (
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

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

func getTab(options protocol.FormattingOptions) string {
	tabSize := options["tabSize"].(float64)
	insertSpace := options["insertSpaces"].(bool)

	if insertSpace {
		return strings.Repeat(" ", int(tabSize))
	} else {
		return "\t"
	}
}
