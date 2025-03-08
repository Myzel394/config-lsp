package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/utils"
	"maps"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetSectionBodyCompletions(
	d *wireguard.WGDocument,
	section ast.WGSection,
	property *ast.WGProperty,
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
	section ast.WGSection,
	property *ast.WGProperty,
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

	if property == nil || property.Separator == nil {
		// First scenario
		return getKeyCompletions(section), nil
	}

	// Check if the cursor it outside the value
	position := common.LSPCharacterAsCursorPosition(params.Position.Character)
	if property.Value != nil && property.Value.IsPositionAfterEnd(position) {
		// Then we don't show anything
		return nil, nil
	}

	// Otherwise, suggest value completions
	return getValueCompletions(section, *property, position), nil
}

func getKeyCompletions(
	section ast.WGSection,
) []protocol.CompletionItem {
	options := make(map[string]docvalues.DocumentationValue)

	switch section.Header.Name {
	case "Interface":
		maps.Copy(options, fields.InterfaceOptions)

		// Remove existing, non-duplicate options
		for _, property := range section.Properties {
			if _, found := fields.InterfaceAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			delete(options, property.Key.Name)
		}
	case "Peer":
		maps.Copy(options, fields.PeerOptions)

		// Remove existing, non-duplicate options
		for _, property := range section.Properties {
			if _, found := fields.PeerAllowedDuplicateFields[property.Key.Name]; found {
				continue
			}

			delete(options, property.Key.Name)
		}
	}

	kind := protocol.CompletionItemKindField

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
	)
}

func getValueCompletions(
	section ast.WGSection,
	property ast.WGProperty,
	cursorPosition common.CursorPosition,
) []protocol.CompletionItem {
	// TODO: Normalize section header name
	options, found := fields.OptionsHeaderMap[section.Header.Name]

	if !found {
		return nil
	}

	option, found := options[property.Key.Name]

	if !found {
		return nil
	}

	if property.Value == nil {
		return option.DeprecatedFetchCompletions("", 0)
	} else {
		return option.DeprecatedFetchCompletions(
			property.Value.Value,
			common.DeprecatedImprovedCursorToIndex(
				cursorPosition,
				property.Value.Value,
				property.Value.Start.Character,
			),
		)
	}
}
