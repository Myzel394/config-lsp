package roothandler

import (
	wireguard "config-lsp/handlers/wireguard/lsp"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) (any, error) {
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
	case LanguageFstab:
		fallthrough
	case LanguageHosts:
		fallthrough
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageWireguard:
		return wireguard.TextDocumentCodeAction(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
