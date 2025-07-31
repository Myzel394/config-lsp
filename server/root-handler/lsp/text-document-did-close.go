package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	bitcoinconf "config-lsp/handlers/bitcoin_conf/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	document := shared.GetDocument(params.TextDocument.URI)

	if document == nil {
		return nil
	}

	shared.RemoveDocument(params.TextDocument.URI)

	switch *document.Language {
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidClose(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentDidClose(context, params)
	case shared.LanguageFstab:
		return fstab.TextDocumentDidClose(context, params)
	case shared.LanguageWireguard:
		return wireguard.TextDocumentDidClose(context, params)
	case shared.LanguageHosts:
		return hosts.TextDocumentDidClose(context, params)
	case shared.LanguageAliases:
		return aliases.TextDocumentDidClose(context, params)
	case shared.LanguageBitcoinConf:
		return bitcoinconf.TextDocumentDidClose(context, params)
	default:
	}

	return nil
}
