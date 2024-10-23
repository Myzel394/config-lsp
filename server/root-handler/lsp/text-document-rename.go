package lsp

import (
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
		return nil, utils.LanguageUndetectableError{}
	}

	switch *language {
	case utils.LanguageHosts:
		return nil, nil
	case utils.LanguageSSHDConfig:
		return nil, nil
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentRename(context, params)
	case utils.LanguageFstab:
		return nil, nil
	case utils.LanguageWireguard:
		return nil, nil
	case utils.LanguageAliases:
		return aliases.TextDocumentRename(context, params)
	}

	panic("root-handler/TextDocumentRename: unexpected language" + *language)
}
