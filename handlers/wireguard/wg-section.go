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
	for _, property := range s.Properties {
		if property.Key.Location.Start <= lineNumber && property.Key.Location.End >= lineNumber {
			return &property, nil
		}
	}

	return nil, propertyNotFoundError{}
}

func (s wireguardSection) getCompletionsForEmptyLine() ([]protocol.CompletionItem, error) {
	if s.Name == nil {
		return nil, nil
	}

	switch *s.Name {
	case "Interface":
		availableOptions := map[docvalues.EnumString]docvalues.Value{}

		maps.Copy(availableOptions, interfaceOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := interfaceAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			// Remove the option from the available options
			maps.DeleteFunc(
				availableOptions,
				func(key docvalues.EnumString, value docvalues.Value) bool {
					return key.DescriptionText == property.Key.Name
				},
			)
		}

		kind := protocol.CompletionItemKindProperty

		return utils.MapMapToSlice(
			availableOptions,
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

	return []protocol.CompletionItem{}, nil
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
		var insertText string

		if character == property.Key.Location.End {
			insertText = " = "
		} else {
			insertText = "= "
		}

		return []protocol.CompletionItem{
			{
				Label:      "=",
				InsertText: &insertText,
			},
		}, propertyNotFullyTypedError{}
	}

	var option docvalues.Value

	for enum, opt := range interfaceOptions {
		if enum.InsertText == property.Key.Name {
			option = opt
			break
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
