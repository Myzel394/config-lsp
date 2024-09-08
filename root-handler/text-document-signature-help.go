package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
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
		return nil, nil
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageFstab:
		return nil, nil
	case LanguageWireguard:
		return nil, nil
	case LanguageAliases:
		return aliases.TextDocumentSignatureHelp(context, params)
	}

	panic("root-handler/TextDocumentSignatureHelp: unexpected language" + *language)
}
