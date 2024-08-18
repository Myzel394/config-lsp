package wireguard

import (
	"config-lsp/common"
	"regexp"
	"slices"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var commentPattern = regexp.MustCompile(`^\s*(;|#)`)
var emptyLinePattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[`)

type characterLocation struct {
	Start uint32
	End   uint32
}

type wireguardLineIndex struct {
	Type             lineType
	BelongingSection *wireguardSection
}

type wireguardParser struct {
	// <key = name>: if nil then does not belong to a section
	Sections []*wireguardSection
	// Used to identify where not to show diagnostics
	CommentLines map[uint32]struct{}

	// Indexes
	LineIndexes map[uint32]wireguardLineIndex
}

func (p *wireguardParser) getSectionByLine(line uint32) *wireguardSection {
	for _, section := range p.Sections {
		if section.StartLine <= line && section.EndLine >= line {
			return section
		}
	}

	return nil
}

// Search for a property by name
// Returns (line number, property)
func (p *wireguardParser) fetchPropertyByName(name string) (*uint32, *wireguardProperty) {
	for _, section := range p.Sections {
		for lineNumber, property := range section.Properties {
			if property.Key.Name == name {
				return &lineNumber, &property
			}
		}
	}

	return nil, nil
}

func (p *wireguardParser) clear() {
	p.Sections = []*wireguardSection{}
	p.CommentLines = map[uint32]struct{}{}
	p.LineIndexes = map[uint32]wireguardLineIndex{}
}

func (p wireguardParser) getInterfaceSection() (*wireguardSection, bool) {
	for _, section := range p.Sections {
		if section.Name != nil && *section.Name == "Interface" {
			return section, true
		}
	}

	return nil, false
}

func getHeaderCompletion(name string, documentation string) protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindEnum

	insertText := "[" + name + "]\n"

	return protocol.CompletionItem{
		Label:            "[" + name + "]",
		InsertTextFormat: &textFormat,
		InsertText:       &insertText,
		Kind:             &kind,
		Documentation:    &documentation,
	}
}

func (p wireguardParser) getRootCompletionsForEmptyLine() []protocol.CompletionItem {
	completions := []protocol.CompletionItem{}

	if _, found := p.getInterfaceSection(); !found {
		completions = append(completions, getHeaderCompletion("Interface", headerInterfaceEnum.Documentation))
	}

	completions = append(completions, getHeaderCompletion("Peer", headerPeerEnum.Documentation))

	return completions
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
	var errors []common.ParseError
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
			p.CommentLines[currentLineNumber] = struct{}{}
			p.LineIndexes[currentLineNumber] = wireguardLineIndex{
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

			section := createWireguardSection(
				currentLineNumber,
				lastLine,
				line,
				collectedProperties,
			)

			p.Sections = append(p.Sections, &section)

			// Add indexes
			for lineNumber := range collectedProperties {
				p.LineIndexes[lineNumber] = wireguardLineIndex{
					Type:             LineTypeProperty,
					BelongingSection: &section,
				}
			}
			p.LineIndexes[currentLineNumber] = wireguardLineIndex{
				Type:             LineTypeHeader,
				BelongingSection: &section,
			}

			// Reset
			collectedProperties = wireguardProperties{}
			lastPropertyLine = nil
		}
	}

	var emptySection *wireguardSection

	if len(collectedProperties) > 0 {
		var endLine uint32

		if len(p.Sections) == 0 {
			endLine = uint32(len(lines))
		} else {
			endLine = p.Sections[len(p.Sections)-1].StartLine
		}

		emptySection = &wireguardSection{
			StartLine:  0,
			EndLine:    endLine,
			Properties: collectedProperties,
		}

		p.Sections = append(p.Sections, emptySection)

		for lineNumber := range collectedProperties {
			p.LineIndexes[lineNumber] = wireguardLineIndex{
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
			emptySection = &wireguardSection{
				StartLine:  0,
				EndLine:    endLine,
				Properties: collectedProperties,
			}

			p.Sections = append(p.Sections, emptySection)

			for newLine := uint32(0); newLine < endLine; newLine++ {
				if _, found := p.LineIndexes[newLine]; found {
					continue
				}

				p.LineIndexes[newLine] = wireguardLineIndex{
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
			if _, found := p.LineIndexes[newLine]; found {
				continue
			}

			p.LineIndexes[newLine] = wireguardLineIndex{
				Type:             LineTypeEmpty,
				BelongingSection: section,
			}
		}
	}

	return errors
}

func (p wireguardParser) getTypeByLine(line uint32) lineType {
	// Check if line is a comment
	if _, found := p.CommentLines[line]; found {
		return LineTypeComment
	}

	if info, found := p.LineIndexes[line]; found {
		return info.Type
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
func (p *wireguardParser) getBelongingSectionByLine(line uint32) *wireguardSection {
	if info, found := p.LineIndexes[line]; found {
		return info.BelongingSection
	}

	return nil
}

func (p *wireguardParser) getPropertyByLine(line uint32) (*wireguardSection, *wireguardProperty) {
	section := p.getSectionByLine(line)

	if section == nil || section.Name == nil {
		return nil, nil
	}

	property, _ := section.findProperty(line)

	if property == nil {
		return nil, nil
	}

	return section, property
}
