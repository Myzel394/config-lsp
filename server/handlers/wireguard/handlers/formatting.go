package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/parsers/ini"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FormatDocument(
	d *wireguard.WGDocument,
	textRange protocol.Range,
	options protocol.FormattingOptions,
) ([]protocol.TextEdit, error) {
	edits := make([]protocol.TextEdit, 0)

	for _, section := range d.Config.Sections {
		it := section.Properties.Iterator()

		for it.Next() {
			property := it.Value().(*ini.Property)

			newEdits, err := ini.FormatProperty(property, options)

			if err == nil {
				edits = append(edits, newEdits...)
			}
		}
	}

	return edits, nil
}
