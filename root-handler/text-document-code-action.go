package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
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
		return nil, nil
	case LanguageHosts:
		return hosts.TextDocumentCodeAction(context, params)
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageWireguard:
		return wireguard.TextDocumentCodeAction(context, params)
	case LanguageAliases:
		return aliases.TextDocumentCodeAction(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
