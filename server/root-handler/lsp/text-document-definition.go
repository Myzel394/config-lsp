package lsp

import (
	"config-lsp/common"
	aliases "config-lsp/handlers/aliases/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	document := shared.GetDocument(params.TextDocument.URI)

	if document == nil {
		if common.ServerOptions.NoUndetectableErrors {
			return nil, nil
		} else {
			return nil, utils.LanguageUndetectableError{}
		}
	}

	switch *document.Language {
	case shared.LanguageHosts:
		return nil, nil
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDefinition(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentDefinition(context, params)
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageWireguard:
		return nil, nil
	case shared.LanguageAliases:
		return aliases.TextDocumentDefinition(context, params)
	case shared.LanguageBitcoinConf:
		return nil, nil
	}

	panic("root-handler/TextDocumentDefinition: unexpected language" + *document.Language)
}
