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

	// Special case, key defined but separator missing
	if property != nil && property.Separator == nil && !property.Key.ContainsPosition(position) {
		return getKeyCompletions(section, true)
	}

	// First scenario, user adds a new property
	if property == nil || property.Key.Name == "" || property.Key.ContainsPosition(position) {
		return getKeyCompletions(section, false)
	}

	// Check if the cursor it outside the value
	if property.Value != nil && property.Value.IsPositionAfterEnd(position) {
		// Then we don't show anything
		return nil, nil
	}

	// Otherwise, suggest value completions
	return getValueCompletions(section, property, position)
}

func getKeyCompletions(
	currentSection *ini.Section,
	onlySuggestSeparator bool,
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

	kind := protocol.CompletionItemKindField

	return utils.MapMapToSlice(
		options,
		func(optionName string, value docvalues.DocumentationValue) protocol.CompletionItem {
			var label string
			var insertText string

			if onlySuggestSeparator {
				label = optionName + " = "
				insertText = "= "
			} else {
				label = optionName
				insertText = optionName + " = "
			}

			return protocol.CompletionItem{
				Kind:          &kind,
				Documentation: value.Documentation,
				Label:         label,
				InsertText:    &insertText,
			}
		},
	), nil
}

func getValueCompletions(
	currentSection *ini.Section,
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
