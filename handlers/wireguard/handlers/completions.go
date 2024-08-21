package handlers

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/parser"
	"config-lsp/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"maps"
)

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

func GetRootCompletionsForEmptyLine(
	p parser.WireguardParser,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	if _, found := p.GetInterfaceSection(); !found {
		completions = append(completions, getHeaderCompletion("Interface", fields.HeaderInterfaceEnum.Documentation))
	}

	completions = append(completions, getHeaderCompletion("Peer", fields.HeaderPeerEnum.Documentation))

	return completions, nil
}

func GetCompletionsForSectionEmptyLine(
	s parser.WireguardSection,
) ([]protocol.CompletionItem, error) {
	if s.Name == nil {
		return nil, nil
	}

	options := make(map[string]docvalues.DocumentationValue)

	switch *s.Name {
	case "Interface":
		maps.Copy(options, fields.InterfaceOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := fields.InterfaceAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			delete(options, property.Key.Name)
		}
	case "Peer":
		maps.Copy(options, fields.PeerOptions)

		// Remove existing options
		for _, property := range s.Properties {
			if _, found := fields.PeerAllowedDuplicateFields[property.Key.Name]; found {
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

func GetSeparatorCompletion(property parser.WireguardProperty, character uint32) ([]protocol.CompletionItem, error) {
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
	}, parser.PropertyNotFullyTypedError{}
}

func GetCompletionsForSectionPropertyLine(
	s parser.WireguardSection,
	lineNumber uint32,
	character uint32,
) ([]protocol.CompletionItem, error) {
	property, err := s.GetPropertyByLine(lineNumber)

	if err != nil {
		return nil, err
	}

	if s.Name == nil {
		return nil, parser.PropertyNotFoundError{}
	}

	options, found := fields.OptionsHeaderMap[*s.Name]

	if !found {
		return nil, parser.PropertyNotFoundError{}
	}

	if property.Separator == nil {
		if _, found := options[property.Key.Name]; found {
			return GetSeparatorCompletion(*property, character)
		}
		// Get empty line completions
		return nil, parser.PropertyNotFullyTypedError{}
	}

	option, found := options[property.Key.Name]

	if !found {
		if character < property.Separator.Location.Start {
			return nil, parser.PropertyNotFullyTypedError{}
		} else {
			return nil, parser.PropertyNotFoundError{}
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
