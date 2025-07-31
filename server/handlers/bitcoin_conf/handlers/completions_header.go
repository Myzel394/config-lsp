package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"maps"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetSectionHeaderCompletions(
	d *bitcoinconf.BTCDocument,
	existingHeader *ini.Header,
) ([]protocol.CompletionItem, error) {
	chains := make(map[string]string)
	maps.Copy(chains, fields.AvailableSections)

	for _, section := range d.Config.Sections {
		if section.Header == nil {
			// Skip
			continue
		}

		name := section.Header.Name

		delete(chains, name)
	}

	return utils.MapMapToSlice(chains, func(name string, documentation string) protocol.CompletionItem {
		return getHeaderCompletion(name, documentation, existingHeader)
	}), nil
}

func getHeaderCompletion(
	name string,
	documentation string,
	existingHeader *ini.Header,
) protocol.CompletionItem {
	kind := protocol.CompletionItemKindEnum

	if existingHeader == nil {
		text := "[" + name + "]\n"
		return protocol.CompletionItem{
			InsertText:    &text,
			Documentation: &documentation,
			Kind:          &kind,
		}
	} else {
		textFormat := protocol.InsertTextFormatSnippet

		return protocol.CompletionItem{
			Label: "[" + name + "]",
			TextEdit: protocol.TextEdit{
				Range:   existingHeader.ToLSPRange(),
				NewText: "[" + name + "]\n",
			},
			InsertTextFormat: &textFormat,
			Kind:             &kind,
			Documentation:    &documentation,
		}
	}
}
