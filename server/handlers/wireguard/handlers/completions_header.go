package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getHeaderCompletion(name string, documentation string) protocol.CompletionItem {
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
}

func GetSectionHeaderCompletions(
	d *wireguard.WGDocument,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	containsInterfaceSection := false

	for _, section := range d.Config.Sections {
		if section.Header.Name == "Interface" {
			containsInterfaceSection = true
			break
		}
	}

	if !containsInterfaceSection {
		completions = append(completions, getHeaderCompletion("Interface", fields.HeaderInterfaceEnum.Documentation))
	}

	completions = append(completions, getHeaderCompletion("Peer", fields.HeaderPeerEnum.Documentation))

	return completions, nil
}
