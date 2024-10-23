package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, utils.LanguageUndetectableError{}
	}

	switch *language {
	case utils.LanguageHosts:
		return nil, nil
	case utils.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDefinition(context, params)
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentDefinition(context, params)
	case utils.LanguageFstab:
		return nil, nil
	case utils.LanguageWireguard:
		return nil, nil
	case utils.LanguageAliases:
		return aliases.TextDocumentDefinition(context, params)
	}

	panic("root-handler/TextDocumentDefinition: unexpected language" + *language)
}
