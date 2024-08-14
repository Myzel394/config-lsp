package wireguard

import (
	"config-lsp/common"
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
	CommentLines map[uint32]struct{}
}

func (p *wireguardParser) clear() {
	p.Sections = []wireguardSection{}
	p.CommentLines = map[uint32]struct{}{}
}

func createWireguardParser() wireguardParser {
	parser := wireguardParser{}
	parser.clear()

	return parser
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

func (p *wireguardParser) parseFromString(input string) []common.ParseError {
	errors := []common.ParseError{}
	lines := strings.Split(
		input,
		"\n",
	)

	slices.Reverse(lines)

	collectedProperties := wireguardProperties{}
	var lastPropertyLine *uint32
	var earliestPropertyLine *uint32

	for index, line := range lines {
		currentLineNumber := uint32(len(lines) - index - 1)
		lineType := getLineType(line)

		switch lineType {
		case LineTypeComment:
			p.CommentLines[currentLineNumber] = struct{}{}

		case LineTypeEmpty:
			continue

		case LineTypeProperty:
			err := collectedProperties.AddLine(currentLineNumber, line)

			if err != nil {
				errors = append(errors, common.ParseError{
					Line: currentLineNumber,
					Err:  err,
				})
				continue
			}

			earliestPropertyLine = &currentLineNumber

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

	if len(collectedProperties) > 0 {
		p.Sections = append(p.Sections, wireguardSection{
			StartLine:  *earliestPropertyLine,
			EndLine:    *lastPropertyLine,
			Name:       nil,
			Properties: collectedProperties,
		})
	}

	// Since we parse the content from bottom to top,
	// we need to reverse the order
	slices.Reverse(p.Sections)

	return errors
}

func (p wireguardParser) getTypeByLine(line uint32) lineType {
	// Check if line is a comment
	if _, found := p.CommentLines[line]; found {
		return LineTypeComment
	}

	// Check if line is a section
	for _, section := range p.Sections {
		if section.StartLine <= line && section.EndLine >= line {
			if section.StartLine == line && section.Name != nil {
				return LineTypeHeader
			}

			// Check for properties
			for propertyLineNumber := range section.Properties {
				if propertyLineNumber == line {
					return LineTypeProperty
				}
			}
		}
	}

	return LineTypeEmpty
}

// Get the section that the line belongs to
// Example:
// [Interface]
// Address = 10.0.0.1
//
// <line here>
// [Peer]
//
// This would return the section [Interface]
func (p wireguardParser) getBelongingSectionByLine(line uint32) *wireguardSection {
	for index := range p.Sections {
		section := p.Sections[len(p.Sections)-index-1]

		if section.StartLine <= line && section.Name != nil {
			return &section
		}
	}

	// Global section
	return nil
}
