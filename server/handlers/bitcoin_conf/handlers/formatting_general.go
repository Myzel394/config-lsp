// Format stuff outside of the AST nodes
package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/parsers/ini"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Make sure there are exactly two newlines between sections in the document.
func formatNewlinesBetweenSections(
	d *bitcoinconf.BTCDocument,
	textRange protocol.Range,
) []protocol.TextEdit {
	println(fmt.Sprintf("Formatting newlines between sections in range: %v; Da lange is %d", textRange, len(d.Config.Sections)))
	if len(d.Config.Sections) < 2 {
		return nil
	}

	edits := make([]protocol.TextEdit, 0, len(d.Config.Sections)-1)

	for _, section := range d.Config.Sections {
		// Skip sections outside the specified range
		if !(section.Start.Line >= textRange.Start.Line && section.End.Line <= textRange.End.Line) {
			continue
		}

		rawLastLine, rawLastProperty := section.Properties.Max()
		lastLine := rawLastLine.(uint32)

		diff := section.End.Line - lastLine

		if diff != 2 {
			lastProperty := rawLastProperty.(*ini.Property)
			edits = append(edits, protocol.TextEdit{
				Range: protocol.Range{
					Start: lastProperty.End.ToLSPPosition(),
					End:   section.End.ToLSPPosition(),
				},
				NewText: "\n\n",
			})
		}
	}

	return edits
}
