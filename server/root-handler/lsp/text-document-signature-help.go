package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, nil
	}

	switch *language {
	case shared.LanguageHosts:
		return hosts.TextDocumentSignatureHelp(context, params)
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentSignatureHelp(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentSignatureHelp(context, params)
	case shared.LanguageFstab:
		return fstab.TextDocumentSignatureHelp(context, params)
	case shared.LanguageWireguard:
		return nil, nil
	case shared.LanguageAliases:
		return aliases.TextDocumentSignatureHelp(context, params)
	}

	panic("root-handler/TextDocumentSignatureHelp: unexpected language" + *language)
}
