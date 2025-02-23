package parser

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
	"regexp"
	"slices"
	"strings"
)

var commentPattern = regexp.MustCompile(`^\s*(;|#)`)
var emptyLinePattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[`)

func (p *ast.WGConfig) ParseFromString(input string) []common.ParseError {
	var errors []common.ParseError
	lines := strings.Split(
		input,
		"\n",
	)

	slices.Reverse(lines)

	collectedProperties := WireguardProperties{}
	var lastPropertyLine *uint32

	for index, line := range lines {
		currentLineNumber := uint32(len(lines) - index - 1)
		lineType := getLineType(line)

		switch lineType {
		case LineTypeComment:
			p.commentLines[currentLineNumber] = struct{}{}
			p.linesIndexes[currentLineNumber] = wireguardLineIndex{
				Type:             LineTypeComment,
				BelongingSection: nil,
			}

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

			section := CreateWireguardSection(
				currentLineNumber,
				lastLine,
				line,
				collectedProperties,
			)

			p.Sections = append(p.Sections, &section)

			// Add indexes
			for lineNumber := range collectedProperties {
				p.linesIndexes[lineNumber] = wireguardLineIndex{
					Type:             LineTypeProperty,
					BelongingSection: &section,
				}
			}
			p.linesIndexes[currentLineNumber] = wireguardLineIndex{
				Type:             LineTypeHeader,
				BelongingSection: &section,
			}

			// Reset
			collectedProperties = WireguardProperties{}
			lastPropertyLine = nil
		}
	}

	var emptySection *ast.WGSection

	if len(collectedProperties) > 0 {
		var endLine uint32

		if len(p.Sections) == 0 {
			endLine = uint32(len(lines))
		} else {
			endLine = p.Sections[len(p.Sections)-1].StartLine
		}

		emptySection = &ast.WGSection{
			StartLine:  0,
			EndLine:    endLine,
			Properties: collectedProperties,
		}

		p.Sections = append(p.Sections, emptySection)

		for lineNumber := range collectedProperties {
			p.linesIndexes[lineNumber] = wireguardLineIndex{
				Type:             LineTypeProperty,
				BelongingSection: emptySection,
			}
		}
		p.Sections = append(p.Sections, emptySection)
	} else {
		// Add empty section
		var endLine = uint32(len(lines))

		if len(p.Sections) > 0 {
			endLine = p.Sections[len(p.Sections)-1].StartLine
		}

		// Add empty section
		if endLine != 0 {
			emptySection = &ast.WGSection{
				StartLine:  0,
				EndLine:    endLine,
				Properties: collectedProperties,
			}

			p.Sections = append(p.Sections, emptySection)

			for newLine := uint32(0); newLine < endLine; newLine++ {
				if _, found := p.linesIndexes[newLine]; found {
					continue
				}

				p.linesIndexes[newLine] = wireguardLineIndex{
					Type:             LineTypeEmpty,
					BelongingSection: emptySection,
				}
			}
		}
	}

	// Since we parse the content from bottom to top, we need to reverse the sections
	// so its in correct order
	slices.Reverse(p.Sections)

	// Fill empty lines between sections
	for lineNumber, section := range p.Sections {
		var endLine uint32

		if len(p.Sections) > lineNumber+1 {
			nextSection := p.Sections[lineNumber+1]
			endLine = nextSection.StartLine
		} else {
			endLine = uint32(len(lines))
		}

		for newLine := section.StartLine; newLine < endLine; newLine++ {
			if _, found := p.linesIndexes[newLine]; found {
				continue
			}

			p.linesIndexes[newLine] = wireguardLineIndex{
				Type:             LineTypeEmpty,
				BelongingSection: section,
			}
		}
	}

	return errors
}

func (p *ast.WGConfig) GetSectionByLine(line uint32) *ast.WGSection {
	for _, section := range p.Sections {
		if section.StartLine <= line && section.EndLine >= line {
			return section
		}
	}

	return nil
}

// Search for a property by name
// Returns (line number, property)
func (p *ast.WGConfig) FindFirstPropertyByName(name string) (*uint32, *ast.WGProperty) {
	for _, section := range p.Sections {
		for lineNumber, property := range section.Properties {
			if property.Key.Name == name {
				return &lineNumber, &property
			}
		}
	}

	return nil, nil
}

func (p ast.WGConfig) GetInterfaceSection() (*ast.WGSection, bool) {
	for _, section := range p.Sections {
		if section.Header != nil && *section.Header == "Interface" {
			return section, true
		}
	}

	return nil, false
}

func (p ast.WGConfig) GetTypeByLine(line uint32) LineType {
	// Check if line is a comment
	if _, found := p.commentLines[line]; found {
		return LineTypeComment
	}

	if info, found := p.linesIndexes[line]; found {
		return info.Type
	}

	return LineTypeEmpty
}

func CreateWireguardParser() ast.WGConfig {
	parser := ast.WGConfig{}
	parser.Clear()

	return parser
}

type LineType string

const (
	LineTypeComment  LineType = "comment"
	LineTypeEmpty    LineType = "empty"
	LineTypeHeader   LineType = "header"
	LineTypeProperty LineType = "property"
)

func getLineType(line string) LineType {
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
