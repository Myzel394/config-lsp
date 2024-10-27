package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	utils "config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) (any, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return utils.FetchAddLanguageActions(params.TextDocument.URI)
	}

	switch *language {
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageHosts:
		return hosts.TextDocumentCodeAction(context, params)
	case shared.LanguageSSHDConfig:
		return nil, nil
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentCodeAction(context, params)
	case shared.LanguageWireguard:
		return wireguard.TextDocumentCodeAction(context, params)
	case shared.LanguageAliases:
		return aliases.TextDocumentCodeAction(context, params)
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + *language)
}
