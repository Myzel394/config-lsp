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

type wireguardSection struct {
	StartLine uint32
	EndLine   uint32
	// nil = do not belong to a section
	Name       *string
	Properties wireguardProperties
}

func (s wireguardSection) String() string {
	var name string

	if s.Name == nil {
		name = "//<nil>//"
	} else {
		name = *s.Name
	}

	return fmt.Sprintf("[%s]; %d-%d: %v", name, s.StartLine, s.EndLine, s.Properties)
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

	options := make(map[docvalues.EnumString]docvalues.Value)

	switch *s.Name {
	case "Interface":
		maps.Copy(options, interfaceOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := interfaceAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			// Remove the option from the available options
			maps.DeleteFunc(
				options,
				func(key docvalues.EnumString, value docvalues.Value) bool {
					return key.DescriptionText == property.Key.Name
				},
			)
		}
	case "Peer":
		maps.Copy(options, peerOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := peerAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			// Remove the option from the available options
			maps.DeleteFunc(
				options,
				func(key docvalues.EnumString, value docvalues.Value) bool {
					return key.DescriptionText == property.Key.Name
				},
			)
		}
	}

	kind := protocol.CompletionItemKindProperty

	return utils.MapMapToSlice(
		options,
		func(key docvalues.EnumString, value docvalues.Value) protocol.CompletionItem {
			insertText := key.InsertText + " = "

			return protocol.CompletionItem{
				Label:         key.InsertText,
				InsertText:    &insertText,
				Documentation: key.Documentation,
				Kind:          &kind,
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

	if property.Separator == nil {
		return getSeparatorCompletion(*property, character)
	}

	var option docvalues.Value

	switch *p.Name {
	case "Interface":
		for enum, opt := range interfaceOptions {
			if enum.InsertText == property.Key.Name {
				option = opt
				break
			}
		}
	case "Peer":
		for enum, opt := range peerOptions {
			if enum.InsertText == property.Key.Name {
				option = opt
				break
			}
		}
	}

	if option == nil {
		return nil, propertyNotFoundError{}
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

func createWireguardSection(startLine uint32, endLine uint32, headerLine string, props wireguardProperties) wireguardSection {
	match := validHeaderPattern.FindStringSubmatch(headerLine)

	var header string

	if match == nil {
		// Still typing it
		header = headerLine[1:]
	} else {
		header = match[1]
	}

	return wireguardSection{
		StartLine:  startLine,
		EndLine:    endLine,
		Name:       &header,
		Properties: props,
	}
}
