package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil
	}

	delete(shared.OpenedFiles, params.TextDocument.URI)
	shared.Handler.RemoveDocument(params.TextDocument.URI)

	switch *language {
	case utils.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidClose(context, params)
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentDidClose(context, params)
	case utils.LanguageFstab:
		return fstab.TextDocumentDidClose(context, params)
	case utils.LanguageWireguard:
		return wireguard.TextDocumentDidClose(context, params)
	case utils.LanguageHosts:
		return hosts.TextDocumentDidClose(context, params)
	case utils.LanguageAliases:
		return aliases.TextDocumentDidClose(context, params)
	default:
	}

	return nil
}
