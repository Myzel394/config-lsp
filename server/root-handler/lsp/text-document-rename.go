package lsp

import (
	"config-lsp/common"
	aliases "config-lsp/handlers/aliases/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRename(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		if common.ServerOptions.NoUndetectableErrors {
			return nil, nil
		} else {
			return nil, utils.LanguageUndetectableError{}
		}
	}

	switch *language {
	case shared.LanguageHosts:
		return nil, nil
	case shared.LanguageSSHDConfig:
		return nil, nil
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentRename(context, params)
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageWireguard:
		return nil, nil
	case shared.LanguageAliases:
		return aliases.TextDocumentRename(context, params)
	case shared.LanguageBitcoinConf:
		return nil, nil
	}

	panic("root-handler/TextDocumentRename: unexpected language" + *language)
}
