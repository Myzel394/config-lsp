package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
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
		return sshdconfig.TextDocumentCompletion(context, params)
	case LanguageWireguard:
		return wireguard.TextDocumentCompletion(context, params)
	case LanguageHosts:
		return hosts.TextDocumentCompletion(context, params)
	case LanguageAliases:
		return aliases.TextDocumentCompletion(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
