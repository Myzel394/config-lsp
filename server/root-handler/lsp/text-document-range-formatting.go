package lsp

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormattingFunc(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
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
		return sshdconfig.TextDocumentRangeFormatting(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentRangeFormatting(context, params)
	case shared.LanguageFstab:
		return nil, nil
	case shared.LanguageWireguard:
		return wireguard.TextDocumentRangeFormatting(context, params)
	case shared.LanguageAliases:
		return nil, nil
	case shared.LanguageBitcoinConf:
		return bitcoinconf.TextDocumentRangeFormatting(context, params)
	}

	panic("root-handler/TextDocumentRangeFormattingFunc: unexpected language" + *document.Language)
}
