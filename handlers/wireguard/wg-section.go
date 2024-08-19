package wireguard

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"fmt"
	"maps"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type propertyNotFoundError struct{}

func (e propertyNotFoundError) Error() string {
	return "Property not found"
}

type propertyNotFullyTypedError struct{}

func (e propertyNotFullyTypedError) Error() string {
	return "Property not fully typed"
}

type wireguardSectionType uint

const (
	wireguardSectionUnknownType   wireguardSectionType = 0
	wireguardSectionInterfaceType wireguardSectionType = 1
	wireguardSectionPeerType      wireguardSectionType = 2
)

type wireguardSection struct {
	Name       *string
	StartLine  uint32
	EndLine    uint32
	Properties wireguardProperties
}

func (s wireguardSection) getHeaderLineRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      s.StartLine,
			Character: 0,
		},
		End: protocol.Position{
			Line:      s.StartLine,
			Character: 99999999,
		},
	}
}

func (s wireguardSection) getRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      s.StartLine,
			Character: 0,
		},
		End: protocol.Position{
			Line:      s.EndLine,
			Character: 99999999,
		},
	}
}

func (s wireguardSection) String() string {
	var name string

	if s.Name == nil {
		name = "<nil>"
	} else {
		name = *s.Name
	}

	return fmt.Sprintf("[%s]; %d-%d: %v", name, s.StartLine, s.EndLine, s.Properties)
}

func (s *wireguardSection) fetchFirstProperty(name string) (*uint32, *wireguardProperty) {
	for line, property := range s.Properties {
		if property.Key.Name == name {
			return &line, &property
		}
	}

	return nil, nil
}

func (s *wireguardSection) existsProperty(name string) bool {
	_, property := s.fetchFirstProperty(name)

	return property != nil
}

func (s *wireguardSection) findProperty(lineNumber uint32) (*wireguardProperty, error) {
	property, found := s.Properties[lineNumber]

	if !found {
		return nil, propertyNotFoundError{}
	}

	return &property, nil
}

func (s wireguardSection) getCompletionsForEmptyLine() ([]protocol.CompletionItem, error) {
	if s.Name == nil {
		return nil, nil
	}

	options := make(map[string]docvalues.DocumentationValue)

	switch *s.Name {
	case "Interface":
		maps.Copy(options, interfaceOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := interfaceAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			delete(options, property.Key.Name)
		}
	case "Peer":
		maps.Copy(options, peerOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := peerAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			delete(options, property.Key.Name)
		}
	}

	kind := protocol.CompletionItemKindProperty

	return utils.MapMapToSlice(
		options,
		func(optionName string, value docvalues.DocumentationValue) protocol.CompletionItem {
			insertText := optionName + " = "

			return protocol.CompletionItem{
				Kind:          &kind,
				Documentation: value.Documentation,
				Label:         optionName,
				InsertText:    &insertText,
			}
		},
	), nil
}

func getSeparatorCompletion(property wireguardProperty, character uint32) ([]protocol.CompletionItem, error) {
	var insertText string

	if character == property.Key.Location.End {
		insertText = property.Key.Name + " = "
	} else {
		insertText = "= "
	}

	kind := protocol.CompletionItemKindValue

	return []protocol.CompletionItem{
		{
			Label:      insertText,
			InsertText: &insertText,
			Kind:       &kind,
		},
	}, propertyNotFullyTypedError{}
}

func (p wireguardSection) getCompletionsForPropertyLine(
	lineNumber uint32,
	character uint32,
) ([]protocol.CompletionItem, error) {
	property, err := p.findProperty(lineNumber)

	if err != nil {
		return nil, err
	}

	if p.Name == nil {
		return nil, propertyNotFoundError{}
	}

	options, found := optionsHeaderMap[*p.Name]

	if !found {
		return nil, propertyNotFoundError{}
	}

	if property.Separator == nil {
		if _, found := options[property.Key.Name]; found {
			return getSeparatorCompletion(*property, character)
		}
		// Get empty line completions
		return nil, propertyNotFullyTypedError{}
	}

	option, found := options[property.Key.Name]

	if !found {
		if character < property.Separator.Location.Start {
			return nil, propertyNotFullyTypedError{}
		} else {
			return nil, propertyNotFoundError{}
		}
	}

	if property.Value == nil {
		if character >= property.Separator.Location.End {
			return option.FetchCompletions("", 0), nil
		}
	}

	relativeCursor := character - property.Value.Location.Start

	return option.FetchCompletions(property.Value.Value, relativeCursor), nil
}

var validHeaderPattern = regexp.MustCompile(`^\s*\[(?P<header>.+?)\]\s*$`)

// Create a new create section
// Return (<name>, <new section>)
func createWireguardSection(
	startLine uint32,
	endLine uint32,
	headerLine string,
	props wireguardProperties,
) wireguardSection {
	match := validHeaderPattern.FindStringSubmatch(headerLine)

	var header string

	if match == nil {
		// Still typing it
		header = headerLine[1:]
	} else {
		header = match[1]
	}

	return wireguardSection{
		Name:       &header,
		StartLine:  startLine,
		EndLine:    endLine,
		Properties: props,
	}
}
