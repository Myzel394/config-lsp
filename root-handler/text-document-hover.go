package roothandler

import (
	"config-lsp/handlers/fstab"
	"config-lsp/handlers/wireguard"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentHover(context, params)
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageWireguard:
		return wireguard.TextDocumentHover(context, params)
	}

	panic("root-handler/TextDocumentHover: unexpected language" + language)
}
