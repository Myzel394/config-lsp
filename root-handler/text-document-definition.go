package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
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
		return sshdconfig.TextDocumentDefinition(context, params)
	case LanguageFstab:
		return nil, nil
	case LanguageWireguard:
		return nil, nil
	case LanguageAliases:
		return aliases.TextDocumentDefinition(context, params)
	}

	panic("root-handler/TextDocumentDefinition: unexpected language" + *language)
}
