package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentPrepareRename(context *glsp.Context, params *protocol.PrepareRenameParams) (any, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, utils.LanguageUndetectableError{}
	}

	switch *language {
	case shared.LanguageHosts:
		return nil, nil
	case shared.LanguageSSHDConfig:
		return nil, nil
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentPrepareRename(context, params)
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageWireguard:
		return nil, nil
	case shared.LanguageAliases:
		return aliases.TextDocumentPrepareRename(context, params)
	}

	panic("root-handler/TextDocumentPrepareRename: unexpected language" + *language)
}
