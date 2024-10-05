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
		return hosts.TextDocumentHover(context, params)
	case LanguageSSHDConfig:
		return sshdconfig.TextDocumentHover(context, params)
	case LanguageSSHConfig:
		return sshconfig.TextDocumentHover(context, params)
	case LanguageFstab:
		return fstab.TextDocumentHover(context, params)
	case LanguageWireguard:
		return wireguard.TextDocumentHover(context, params)
	case LanguageAliases:
		return aliases.TextDocumentHover(context, params)
	}

	panic("root-handler/TextDocumentHover: unexpected language" + *language)
}
