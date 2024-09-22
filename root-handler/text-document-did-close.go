package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		showParseError(
			context,
			params.TextDocument.URI,
			undetectableError,
		)

		return undetectableError.Err
	}

	delete(openedFiles, params.TextDocument.URI)
	rootHandler.RemoveDocument(params.TextDocument.URI)

	switch *language {
	case LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidClose(context, params)
	case LanguageSSHConfig:
		return sshconfig.TextDocumentDidClose(context, params)
	case LanguageFstab:
		return fstab.TextDocumentDidClose(context, params)
	case LanguageWireguard:
		return wireguard.TextDocumentDidClose(context, params)
	case LanguageHosts:
		return hosts.TextDocumentDidClose(context, params)
	case LanguageAliases:
		return aliases.TextDocumentDidClose(context, params)
	default:
	}

	return nil
}
