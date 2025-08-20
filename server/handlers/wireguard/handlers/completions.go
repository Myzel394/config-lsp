package handlers

import (
	"config-lsp/handlers/wireguard"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func SuggestCompletions(
	d *wireguard.WGDocument,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	lineNumber := params.Position.Line

	if _, found := d.Config.CommentLines[lineNumber]; found {
		return nil, nil
	}

	section := d.Config.FindSectionByLine(lineNumber)
	property := d.Config.FindPropertyByLine(lineNumber)

	println(fmt.Sprintf("SuggestCompletions: section: %v, property: %v", section, property))

	if section == nil {
		// First, the user needs to define a section header
		if property == nil {
			print("asdasd")
			return GetSectionHeaderCompletions(d, nil)
		} else {
			// However, if they start typing a property - we should not
			// show anything to signal them that they can't write a property yet.
			return nil, nil
		}
	} else {
		return GetSectionBodyCompletions(d, *section, property, params)
	}
}
