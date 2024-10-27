package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, nil
	}

	switch *language {
	case shared.LanguageHosts:
		return hosts.TextDocumentHover(context, params)
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentHover(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentHover(context, params)
	case shared.LanguageFstab:
		return fstab.TextDocumentHover(context, params)
	case shared.LanguageWireguard:
		return wireguard.TextDocumentHover(context, params)
	case shared.LanguageAliases:
		return aliases.TextDocumentHover(context, params)
	}

	panic("root-handler/TextDocumentHover: unexpected language" + *language)
}
