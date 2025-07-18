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

	entries := d.Config.GetOptionsInRange(textRange.Start.Line, textRange.End.Line)

	for _, info := range entries {
		option := info.Option

		if option.Key == nil {
			continue
		}

		edits = append(edits, formatOption(option, info.Block, options)...)
	}

	return edits, nil
}
