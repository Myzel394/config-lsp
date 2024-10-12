package parser

import (
	"regexp"
	"strings"
)

type ParseFeatures struct {
	ParseDoubleQuotes      bool
	ParseEscapedCharacters bool
	TrimWhitespace         bool
	Replacements           *map[string]string
}

var FullFeatures = ParseFeatures{
	ParseDoubleQuotes:      true,
	ParseEscapedCharacters: true,
	TrimWhitespace:         false,
	Replacements:           &map[string]string{},
}

type ParsedString struct {
	Raw   string
	Value string
}

func ParseRawString(
	raw string,
	features ParseFeatures,
) ParsedString {
	value := raw

	if len(*features.Replacements) > 0 {
		value = ParseReplacements(value, *features.Replacements)
	}

	if features.TrimWhitespace {
		value = TrimWhitespace(value, features.ParseDoubleQuotes)
	}

	// Parse double quotes
	if features.ParseDoubleQuotes {
		value = ParseDoubleQuotes(value)
	}

	// Parse escaped characters
	if features.ParseEscapedCharacters {
		value = ParseEscapedCharacters(value)
	}

	return ParsedString{
		Raw:   raw,
		Value: value,
	}
}

var trimPattern = regexp.MustCompile(`\s+`)

func TrimWhitespace(
	raw string,
	respectDoubleQuotes bool,
) string {
	if !respectDoubleQuotes {
		return trimPattern.ReplaceAllString(
			strings.TrimSpace(raw),
			" ",
		)
	}

	value := raw
	currentIndex := 0

	for {
		nextStart, found := findNextDoubleQuote(value, currentIndex)

		if found {
			part := value[:nextStart]
			value = strings.TrimSpace(part) + value[nextStart:]
		}

		nextEnd, found := findNextDoubleQuote(value, nextStart+1)

		if !found {
			break
		}

		currentIndex = nextEnd + 1
	}

	// last part
	if currentIndex < len(value) {
		part := value[currentIndex:]

		value = value[:currentIndex] + strings.TrimSpace(part)
	}

	return value
}

func ParseDoubleQuotes(
	raw string,
) string {
	value := raw
	currentIndex := 0

	for {
		start, found := findNextDoubleQuote(value, currentIndex)

		if found && start < (len(value)-1) {
			currentIndex = max(0, start-1)
			end, found := findNextDoubleQuote(value, start+1)

			if found {
				insideContent := value[start+1 : end]
				value = modifyString(value, start, end+1, insideContent)

				continue
			}
		}

		break
	}

	return value
}

func ParseEscapedCharacters(
	raw string,
) string {
	value := raw
	currentIndex := 0

	for {
		position, found := findNextEscapedCharacter(value, currentIndex)

		if found {
			currentIndex = max(0, position-1)
			escapedCharacter := value[position+1]
			value = modifyString(value, position, position+2, string(escapedCharacter))
		} else {
			break
		}
	}

	return value
}

func ParseReplacements(
	raw string,
	replacements map[string]string,
) string {
	value := raw

	for key, replacement := range replacements {
		value = strings.ReplaceAll(value, key, replacement)
	}

	return value
}

func modifyString(
	input string,
	start int,
	end int,
	newValue string,
) string {
	return input[:start] + newValue + input[end:]
}

// Find the next non-escaped double quote in [raw] starting from [startIndex]
// When no double quote is found, return -1
// Return as the second argument whether a double quote was found
func findNextDoubleQuote(
	raw string,
	startIndex int,
) (int, bool) {
	for index := startIndex; index < len(raw); index++ {
		if raw[index] == '"' {
			if index == 0 || raw[index-1] != '\\' {
				return index, true
			}
		}
	}

	return -1, false
}

func findNextEscapedCharacter(
	raw string,
	startIndex int,
) (int, bool) {
	for index := startIndex; index < len(raw); index++ {
		if raw[index] == '\\' && index < len(raw)-1 {
			return index, true
		}
	}

	return -1, false
}
