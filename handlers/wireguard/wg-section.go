package wireguard

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"fmt"
	"maps"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

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
