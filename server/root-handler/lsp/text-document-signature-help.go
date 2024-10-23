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

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, nil
	}

	switch *language {
	case utils.LanguageHosts:
		return nil, nil
	case utils.LanguageSSHDConfig:
		return sshdconfig.TextDocumentSignatureHelp(context, params)
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentSignatureHelp(context, params)
	case utils.LanguageFstab:
		return nil, nil
	case utils.LanguageWireguard:
		return nil, nil
	case utils.LanguageAliases:
		return aliases.TextDocumentSignatureHelp(context, params)
	}

	panic("root-handler/TextDocumentSignatureHelp: unexpected language" + *language)
}
