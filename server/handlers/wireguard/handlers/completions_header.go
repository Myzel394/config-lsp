package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

const END_BRACKET = ']'

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

func GetSectionHeaderCompletions(
	d *wireguard.WGDocument,
	existingHeader *ini.Header,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	containsInterfaceSection := d.Config.IncludesHeader("Interface")

	if !containsInterfaceSection {
		completions = append(completions, getHeaderCompletion("Interface", fields.HeaderInterfaceEnum.Documentation, existingHeader))
	}

	completions = append(completions, getHeaderCompletion("Peer", fields.HeaderPeerEnum.Documentation, existingHeader))

	print(fmt.Sprintf("laaaaaaaaaaaaaaaaa appended completions: %v", completions))

	return completions, nil
}
