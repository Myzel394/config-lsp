package roothandler

import (
	aliases "config-lsp/handlers/aliases/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRename(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		showParseError(
			context,
			params.TextDocument.URI,
			undetectableError,
		)

		return nil, undetectableError.Err
	}

	switch *language {
	case LanguageHosts:
		return nil, nil
	case LanguageSSHDConfig:
		return nil, nil
	case LanguageSSHConfig:
		return sshconfig.TextDocumentRename(context, params)
	case LanguageFstab:
		return nil, nil
	case LanguageWireguard:
		return nil, nil
	case LanguageAliases:
		return aliases.TextDocumentRename(context, params)
	}

	panic("root-handler/TextDocumentRename: unexpected language" + *language)
}
