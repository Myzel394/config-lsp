package wireguard

import (
	"regexp"
	"slices"
	"strings"
)

var commentPattern = regexp.MustCompile(`^\s*(;|#)`)
var emptyLinePattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[`)

type characterLocation struct {
	Start uint32
	End   uint32
}

type wireguardParser struct {
	Sections []wireguardSection
	// Used to identify where not to show diagnostics
	CommentLines []uint32
}

type lineType string

const (
	LineTypeComment  lineType = "comment"
	LineTypeEmpty    lineType = "empty"
	LineTypeHeader   lineType = "header"
	LineTypeProperty lineType = "property"
)

func getLineType(line string) lineType {
	if commentPattern.MatchString(line) {
		return LineTypeComment
	}

	if emptyLinePattern.MatchString(line) {
		return LineTypeEmpty
	}

	if headerPattern.MatchString(line) {
		return LineTypeHeader
	}

	return LineTypeProperty
}

func (p *wireguardParser) parseFromString(input string) []lineError {
	errors := []lineError{}
	lines := strings.Split(
		input,
		"\n",
	)

	slices.Reverse(lines)

	collectedProperties := wireguardProperties{}
	var lastPropertyLine *uint32

	for index, line := range lines {
		currentLineNumber := uint32(len(lines) - index - 1)
		lineType := getLineType(line)

		switch lineType {
		case LineTypeComment:
			p.CommentLines = append(p.CommentLines, currentLineNumber)

		case LineTypeEmpty:
			continue

		case LineTypeProperty:
			err := collectedProperties.AddLine(currentLineNumber, line)

			if err != nil {
				errors = append(errors, lineError{
					Line: currentLineNumber,
					Err:  err,
				})
				continue
			}

			if lastPropertyLine == nil {
				lastPropertyLine = &currentLineNumber
			}

		case LineTypeHeader:
			var lastLine uint32

			if lastPropertyLine == nil {
				// Current line
				lastLine = currentLineNumber
			} else {
				lastLine = *lastPropertyLine
			}

			section := createWireguardSection(
				currentLineNumber,
				lastLine,
				line,
				collectedProperties,
			)
			p.Sections = append(p.Sections, section)

			// Reset
			collectedProperties = wireguardProperties{}
			lastPropertyLine = nil
		}
	}

	// Since we parse the content from bottom to top,
	// we need to reverse the order
	slices.Reverse(p.CommentLines)
	slices.Reverse(p.Sections)

	return errors
}
