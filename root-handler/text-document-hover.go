package roothandler

import (
	"config-lsp/handlers/fstab"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		showParseError(
			context,
			params.TextDocument.URI,
			undetectableError,
		)

		return nil, undetectableError.Err
	}

	switch *language {
	case LanguageHosts:
		fallthrough
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageFstab:
		return fstab.TextDocumentHover(context, params)
	case LanguageWireguard:
		return wireguard.TextDocumentHover(context, params)
	}

	panic("root-handler/TextDocumentHover: unexpected language" + *language)
}
