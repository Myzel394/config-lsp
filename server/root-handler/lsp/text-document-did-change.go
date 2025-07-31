package lsp

import (
	"config-lsp/common"
	aliases "config-lsp/handlers/aliases/lsp"
	bitcoinconf "config-lsp/handlers/bitcoin_conf/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	document := shared.GetDocument(params.TextDocument.URI)
	content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text

	// Document not initialized yet
	if document == nil {
		params := &protocol.DidOpenTextDocumentParams{
			TextDocument: protocol.TextDocumentItem{
				URI:     params.TextDocument.URI,
				Text:    content,
				Version: params.TextDocument.Version,
			},
		}
		return TextDocumentDidOpen(context, params)
	}

	newLanguage, err := utils.DetectLanguage(content, "", params.TextDocument.URI)

	// User changed the language OR
	if err == nil && newLanguage != *document.Language {
		newLanguage, err := initFile(
			context,
			content,
			params.TextDocument.URI,
			"",
		)

		if err != nil {
			if common.ServerOptions.NoUndetectableErrors {
				return nil
			} else {
				return err
			}
		}

		params := &protocol.DidOpenTextDocumentParams{
			TextDocument: protocol.TextDocumentItem{
				URI:        params.TextDocument.URI,
				Text:       content,
				Version:    params.TextDocument.Version,
				LanguageID: string(*newLanguage),
			},
		}

		return TextDocumentDidOpen(context, params)
	}

	switch *document.Language {
	case shared.LanguageFstab:
		return fstab.TextDocumentDidChange(context, params)
	case shared.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidChange(context, params)
	case shared.LanguageSSHConfig:
		return sshconfig.TextDocumentDidChange(context, params)
	case shared.LanguageWireguard:
		return wireguard.TextDocumentDidChange(context, params)
	case shared.LanguageHosts:
		return hosts.TextDocumentDidChange(context, params)
	case shared.LanguageAliases:
		return aliases.TextDocumentDidChange(context, params)
	case shared.LanguageBitcoinConf:
		return bitcoinconf.TextDocumentDidChange(context, params)
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + *document.Language)
}
