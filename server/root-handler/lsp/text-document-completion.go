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

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, nil
	}

	switch *language {
	case shared.LanguageFstab:
		return fstab.TextDocumentCompletion(context, params)
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentCompletion(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentCompletion(context, params)
	case shared.LanguageWireguard:
		return wireguard.TextDocumentCompletion(context, params)
	case shared.LanguageHosts:
		return hosts.TextDocumentCompletion(context, params)
	case shared.LanguageAliases:
		return aliases.TextDocumentCompletion(context, params)
	case shared.LanguageBitcoinConf:
		return bitcoinconf.TextDocumentCompletion(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
