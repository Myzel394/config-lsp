package roothandler

import (
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormattingFunc(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
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
		return sshdconfig.TextDocumentRangeFormatting(context, params)
	case LanguageFstab:
		return nil, nil
	case LanguageWireguard:
		return nil, nil
	case LanguageAliases:
		return nil, nil
	}

	panic("root-handler/TextDocumentRangeFormattingFunc: unexpected language" + *language)
}
