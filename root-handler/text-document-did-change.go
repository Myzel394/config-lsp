package roothandler

import (
	"config-lsp/handlers/fstab"
	hosts "config-lsp/handlers/hosts/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text
		newLanguage, err := initFile(
			context,
			content,
			params.TextDocument.URI,
			"",
		)

		if err != nil {
			return err
		}

		language = newLanguage

		params := &protocol.DidOpenTextDocumentParams{
			TextDocument: protocol.TextDocumentItem{
				URI:        params.TextDocument.URI,
				Text:       content,
				Version:    params.TextDocument.Version,
				LanguageID: string(*language),
			},
		}

		switch *language {
		case LanguageFstab:
			return fstab.TextDocumentDidOpen(context, params)
		case LanguageSSHDConfig:
			break
		case LanguageWireguard:
			return wireguard.TextDocumentDidOpen(context, params)
		case LanguageHosts:
			return hosts.TextDocumentDidOpen(context, params)
		}
	}

	switch *language {
	case LanguageFstab:
		return fstab.TextDocumentDidChange(context, params)
	case LanguageSSHDConfig:
		return nil
	case LanguageWireguard:
		return wireguard.TextDocumentDidChange(context, params)
	case LanguageHosts:
		return hosts.TextDocumentDidChange(context, params)
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + *language)
}
