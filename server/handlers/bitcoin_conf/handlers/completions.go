package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func SuggestCompletions(
	d *bitcoinconf.BTCDocument,
	params *protocol.CompletionParams,
) ([]protocol.CompletionItem, error) {
	lineNumber := params.Position.Line

	if _, found := d.Config.CommentLines[lineNumber]; found {
		return nil, nil
	}

	section := d.Config.FindSectionByLine(lineNumber)
	property := d.Config.FindPropertyByLine(lineNumber)

	if section.Header != nil && section.Header.Start.Line == lineNumber {
		// If the user is on a section header line, suggest section headers
		return GetSectionHeaderCompletions(d)
	} else {
		return GetPropertyCompletions(
			d,
			section,
			property,
			params,
		)
	}
}
