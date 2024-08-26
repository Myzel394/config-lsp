package roothandler

import (
	"config-lsp/handlers/fstab"
	hosts "config-lsp/handlers/hosts/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
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
		return fstab.TextDocumentCompletion(context, params)
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageWireguard:
		return wireguard.TextDocumentCompletion(context, params)
	case LanguageHosts:
		return hosts.TextDocumentCompletion(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
