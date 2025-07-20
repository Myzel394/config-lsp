package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FormatDocument(
	d *bitcoinconf.BTCDocument,
	textRange protocol.Range,
	options protocol.FormattingOptions,
) ([]protocol.TextEdit, error) {
	edits := make([]protocol.TextEdit, 0)

	entries := d.Config.GetPropertesInRange(textRange.Start.Line, textRange.End.Line)

	for _, info := range entries {
		edits = append(edits, formatProperty(info.Property, options)...)
	}

	edits = append(edits, formatNewlinesBetweenSections(d, textRange)...)
	edits = append(edits, formatRemoveNewlinesAfterHeader(d, textRange)...)

	return edits, nil
}
