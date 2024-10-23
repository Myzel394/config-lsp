package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
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
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

	if language == nil {
		content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text
		newLanguage, err := initFile(
			context,
			content,
			params.TextDocument.URI,
			"",
		)

		if err != nil {
			return err
		}

		language = newLanguage

		params := &protocol.DidOpenTextDocumentParams{
			TextDocument: protocol.TextDocumentItem{
				URI:        params.TextDocument.URI,
				Text:       content,
				Version:    params.TextDocument.Version,
				LanguageID: string(*language),
			},
		}

		switch *language {
		case utils.LanguageFstab:
			return fstab.TextDocumentDidOpen(context, params)
		case utils.LanguageSSHDConfig:
			return sshdconfig.TextDocumentDidOpen(context, params)
		case utils.LanguageSSHConfig:
			return sshconfig.TextDocumentDidOpen(context, params)
		case utils.LanguageWireguard:
			return wireguard.TextDocumentDidOpen(context, params)
		case utils.LanguageHosts:
			return hosts.TextDocumentDidOpen(context, params)
		case utils.LanguageAliases:
			return aliases.TextDocumentDidOpen(context, params)
		}
	}

	switch *language {
	case utils.LanguageFstab:
		return fstab.TextDocumentDidChange(context, params)
	case utils.LanguageSSHDConfig:
		return sshdconfig.TextDocumentDidChange(context, params)
	case utils.LanguageSSHConfig:
		return sshconfig.TextDocumentDidChange(context, params)
	case utils.LanguageWireguard:
		return wireguard.TextDocumentDidChange(context, params)
	case utils.LanguageHosts:
		return hosts.TextDocumentDidChange(context, params)
	case utils.LanguageAliases:
		return aliases.TextDocumentDidChange(context, params)
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + *language)
}
