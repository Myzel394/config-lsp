package lsp

import (
	"config-lsp/common"
	"config-lsp/root-handler/shared"
	"config-lsp/root-handler/utils"
	"fmt"

	aliases "config-lsp/handlers/aliases/lsp"
	fstab "config-lsp/handlers/fstab/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
	sshdconfig "config-lsp/handlers/sshd_config/lsp"
	wireguard "config-lsp/handlers/wireguard/lsp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	// Find the file type
	content := params.TextDocument.Text
	language, err := initFile(
		context,
		content,
		params.TextDocument.URI,
		params.TextDocument.LanguageID,
	)

	if err != nil {
		return err
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

	panic(fmt.Sprintf("unexpected roothandler.SupportedLanguage: %#v", language))
}

func initFile(
	context *glsp.Context,
	content string,
	uri protocol.DocumentUri,
	advertisedLanguage string,
) (*shared.SupportedLanguage, error) {
	language, err := utils.DetectLanguage(content, advertisedLanguage, uri)

	if err != nil {
		utils.NotifyLanguageUndetectable(context, uri)

		return nil, utils.LanguageUndetectableError{}
	} else {
		utils.NotifyDetectedLanguage(context, uri, language)
	}

	shared.OpenedFiles[uri] = struct{}{}

	// Everything okay, now we can handle the file
	shared.Handler.AddDocument(uri, language)

	return &language, nil
}
