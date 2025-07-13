package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"maps"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetSectionBodyCompletions(
	d *wireguard.WGDocument,
	section ini.Section,
	property *ini.Property,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	// These are the possible scenarios:
	// | = Cursor position
	/*
		[Inter|
	*/
	/*
		[Interface]
		|
	*/
	/*
		[Interface]
		Add|
	*/
	/*
		[Interface]
		Address = 10.0.0.1/24

		|
	*/

	// First scenario, user is typing the section name
	if params.Position.Line == section.Start.Line {
		return GetSectionHeaderCompletions(d)
	}

	// Third and fourth scenarios, the user wants to add a new property
	completions, err := getPropertyCompletions(d, section, property, params)

	if err != nil {
		// Something weird happened
		return completions, err
	}

	// Fourth scenario may arrive here, the user is typing a property name, but the previous line is empty.
	// In this case, the user may want to add a property or add a new section.
	// We should therefore suggest both options.

	isLineEmpty := property == nil
	if !isLineEmpty {
		return completions, nil
	}

	// Check if previous line is empty
	previousLineProperty := d.Config.FindPropertyByLine(params.Position.Line - 1)

	if previousLineProperty == nil && params.Position.Line-1 != section.Start.Line {
		sectionCompletions, err := GetSectionHeaderCompletions(d)

		if err != nil {
			return sectionCompletions, err
		}

		completions = append(completions, sectionCompletions...)
	}

	return completions, nil
}

func getPropertyCompletions(
	d *wireguard.WGDocument,
	section ini.Section,
	property *ini.Property,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	// These are the possible scenarios:
	/* Empty line / Key started / Separator missing:
	Add|
	Address |
	*/
	/* Value missing or started:
	Address = 10.|
	*/

	currentLine := params.Position.Line
	position := common.LSPCharacterAsCursorPosition(params.Position.Character)

	// Special case, key defined but separator missing
	if property != nil && property.Separator == nil && !property.Key.ContainsPosition(position) {
		return getKeyCompletions(section, true, currentLine), nil
	}

	if property == nil || property.Separator == nil || property.Key.ContainsPosition(position) {
		// First scenario
		return getKeyCompletions(section, false, currentLine), nil
	}

	// Check if the cursor it outside the value
	if property.Value != nil && property.Value.IsPositionAfterEnd(position) {
		// Then we don't show anything
		return nil, nil
	}

	// Otherwise, suggest value completions
	return getValueCompletions(section, property, position), nil
}

func getKeyCompletions(
	section ini.Section,
	onlySeparator bool,
	currentLine uint32,
) []protocol.CompletionItem {
	options := make(map[fields.NormalizedName]docvalues.DocumentationValue)
	allowedDuplicatedFields := make(map[fields.NormalizedName]struct{})

	switch section.Header.Name {
	case "Interface":
		maps.Copy(options, fields.InterfaceOptions)
		allowedDuplicatedFields = fields.InterfaceAllowedDuplicateFields
	case "Peer":
		maps.Copy(options, fields.PeerOptions)
		allowedDuplicatedFields = fields.PeerAllowedDuplicateFields
	}

	// Remove existing, non-duplicate options
	it := section.Properties.Iterator()
	for it.Next() {
		iniProperty := it.Value().(*ini.Property)
		normalizedName := fields.CreateNormalizedName(iniProperty.Key.Name)
		if _, found := allowedDuplicatedFields[normalizedName]; found {
			continue
		}

		if iniProperty.Key.Start.Line == currentLine {
			// The user is currently typing the key, thus we should suggest it
			continue
		}

		delete(options, normalizedName)
	}

	kind := protocol.CompletionItemKindField

	return utils.MapMapToSlice(
		options,
		func(rawOptionName fields.NormalizedName, value docvalues.DocumentationValue) protocol.CompletionItem {
			optionName := fields.AllOptionsFormatted[rawOptionName]
			var label string
			var insertText string

			if onlySeparator {
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
	)
}

func getValueCompletions(
	section ini.Section,
	property *ini.Property,
	cursor common.CursorPosition,
) []protocol.CompletionItem {
	// TODO: Normalize section header name
	normalizedHeaderName := fields.CreateNormalizedName(section.Header.Name)
	options, found := fields.OptionsHeaderMap[normalizedHeaderName]

	if !found {
		return nil
	}

	option, found := options[fields.CreateNormalizedName(property.Key.Name)]

	if !found {
		return nil
	}

	if property.Value == nil {
		return option.FetchCompletions("", 0)
	} else {
		return option.FetchCompletions(
			property.Value.Value,
			cursor.ShiftHorizontal(-int(property.Value.Start.Character)),
		)
	}
}
