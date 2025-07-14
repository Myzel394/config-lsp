package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/utils"
	"maps"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetSectionHeaderCompletions(
	d *bitcoinconf.BTCDocument,
) ([]protocol.CompletionItem, error) {
	chains := make(map[string]string)
	maps.Copy(chains, fields.AvailableSections)

	for _, section := range d.Config.Sections {
		if section.Header == nil {
			// Skip
			continue
		}

		name := section.Header.Name

		println("GetSectionHeaderCompletions: section header", name)
		if _, found := chains[name]; found {
			// Delete
			delete(chains, name)
		}
	}

	return utils.MapMapToSlice(chains, func(name string, documentation string) protocol.CompletionItem {
		textFormat := protocol.InsertTextFormatPlainText
		kind := protocol.CompletionItemKindEnum

		insertText := "[" + name + "]\n"

		return protocol.CompletionItem{
			Label:            "[" + name + "]",
			InsertTextFormat: &textFormat,
			InsertText:       &insertText,
			Kind:             &kind,
			Documentation:    &documentation,
		}
	}), nil
}
