package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	language := shared.Handler.GetLanguageForDocument(params.TextDocument.URI)

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

	if newLanguage != language {
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
		case shared.LanguageFstab:
			return fstab.TextDocumentDidOpen(context, params)
		case shared.LanguageSSHDConfig:
			return sshdconfig.TextDocumentDidOpen(context, params)
		case shared.LanguageSSHConfig:
			return sshconfig.TextDocumentDidOpen(context, params)
		case shared.LanguageWireguard:
			return wireguard.TextDocumentDidOpen(context, params)
		case shared.LanguageHosts:
			return hosts.TextDocumentDidOpen(context, params)
		case shared.LanguageAliases:
			return aliases.TextDocumentDidOpen(context, params)
		}
	}

	switch *language {
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
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + *language)
}
