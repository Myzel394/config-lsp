package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getHeaderCompletion(
	name string,
	documentation string,
	suggestEndBracket bool,
) protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindEnum

	var insertText string
	if suggestEndBracket {
		insertText = "[" + name + "]\n"
	} else {
		insertText = "[" + name + "\n"
	}

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
	line string,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	containsInterfaceSection := false

	for _, section := range d.Config.Sections {
		if section.Header.Name == "Interface" {
			containsInterfaceSection = true
			break
		}
	}

	containsEndBracket := line != "" && line[len(line)-1] == ']'

	if !containsInterfaceSection {
		completions = append(completions, getHeaderCompletion("Interface", fields.HeaderInterfaceEnum.Documentation, !containsEndBracket))
	}

	completions = append(completions, getHeaderCompletion("Peer", fields.HeaderPeerEnum.Documentation, !containsEndBracket))

	return completions, nil
}
