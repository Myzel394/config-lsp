package roothandler

import (
	"config-lsp/handlers/wireguard"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	delete(openedFiles, params.TextDocument.URI)
	rootHandler.RemoveDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
	case LanguageSSHDConfig:
	case LanguageWireguard:
		return wireguard.TextDocumentDidClose(context, params)
	default:
	}

	return nil
}
