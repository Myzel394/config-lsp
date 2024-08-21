package roothandler

import (
	"config-lsp/handlers/fstab"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentDidChange(context, params)
	case LanguageSSHDConfig:
		return nil
	case LanguageWireguard:
		return wireguard.TextDocumentDidChange(context, params)
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + language)
}
