package lsp

import (
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormattingFunc(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		return nil, utils.LanguageUndetectableError{}
	}

	switch *language {
	case utils.LanguageHosts:
		return nil, nil
	case utils.LanguageSSHDConfig:
		return sshdconfig.TextDocumentRangeFormatting(context, params)
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentRangeFormatting(context, params)
	case utils.LanguageFstab:
		return nil, nil
	case utils.LanguageWireguard:
		return nil, nil
	case utils.LanguageAliases:
		return nil, nil
	}

	panic("root-handler/TextDocumentRangeFormattingFunc: unexpected language" + *language)
}
