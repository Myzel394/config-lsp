package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"maps"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetPropertyCompletions(
	d *bitcoinconf.BTCDocument,
	section *ini.Section,
	property *ini.Property,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	position := common.LSPCharacterAsCursorPosition(params.Position.Character)

	if property == nil || property.Key.ContainsPosition(position) {
		// First scenario
		return getKeyCompletions(section, property, params)
	}

	// Check if the cursor it outside the value
	if property.Value != nil && property.Value.IsPositionAfterEnd(position) {
		// Then we don't show anything
		return nil, nil
	}

	// Otherwise, suggest value completions
	return getValueCompletions(property, position)
}

func getKeyCompletions(
	currentSection *ini.Section,
	property *ini.Property,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	options := make(map[string]docvalues.DocumentationValue)

	maps.Copy(options, fields.Options)

	// Remove already defined properties
	it := currentSection.Properties.Iterator()
	for it.Next() {
		property := it.Value().(*ini.Property)

		if property.Key.Name == "" {
			continue
		}

		if utils.KeyExists(fields.AllowedDuplicateOptions, property.Key.Name) {
			// If the property is allowed to be duplicated, we keep it
			continue
		}

		// Otherwise, remove it from the options
		delete(options, property.Key.Name)
	}

	var start uint32
	var end uint32

	if property == nil {
		start = 0
		end = 0
	} else {
		start = property.Key.Start.Character

		if property.Value != nil {
			end = property.Value.Start.Character
		} else if property.Separator != nil {
			end = property.Separator.End.Character
		} else {
			end = property.Key.End.Character
		}
	}

	kind := protocol.CompletionItemKindField

	return utils.MapMapToSlice(
		options,
		func(optionName string, value docvalues.DocumentationValue) protocol.CompletionItem {
			insertText := optionName + " = "
			insertFormat := protocol.InsertTextFormatSnippet

			return protocol.CompletionItem{
				Label:            optionName,
				Kind:             &kind,
				Documentation:    value.Documentation,
				InsertTextFormat: &insertFormat,
				TextEdit: protocol.TextEdit{
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      params.Position.Line,
							Character: start,
						},
						End: protocol.Position{
							Line:      params.Position.Line,
							Character: end,
						},
					},
					NewText: insertText,
				},
			}
		},
	), nil
}

func getValueCompletions(
	property *ini.Property,
	cursor common.CursorPosition,
) ([]protocol.CompletionItem, error) {
	option, found := fields.Options[property.Key.Name]

	if !found {
		// If the property is not found in the options, we return an empty slice
		return nil, nil
	}

	if property.Value == nil {
		return option.FetchCompletions("", 0), nil
	} else {
		return option.FetchCompletions(
			property.Value.Value,
			cursor.ShiftHorizontal(-int(property.Value.Start.Character)),
		), nil
	}
}
