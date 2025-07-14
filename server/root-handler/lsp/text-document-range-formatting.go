package lsp

import (
	"config-lsp/common"
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
		return sshdconfig.TextDocumentRangeFormatting(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentRangeFormatting(context, params)
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageWireguard:
		return nil, nil
	case shared.LanguageAliases:
		return nil, nil
	case shared.LanguageBitcoinConf:
		return nil, nil
	}

	panic("root-handler/TextDocumentRangeFormattingFunc: unexpected language" + *language)
}
