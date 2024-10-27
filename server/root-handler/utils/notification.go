package utils

import (
	"config-lsp/root-handler/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type lspNotification struct {
	Uri string
}

type lspDetectedLanguage struct {
	lspNotification

	Language string
}

func NotifyLanguageUndetectable(context *glsp.Context, uri protocol.DocumentUri) {
	go context.Notify(
		"$/config-lsp/languageUndetectable",
		lspNotification{
			Uri: string(uri),
		},
	)

	go context.Notify(
		"window/showMessage",
		protocol.ShowMessageParams{
			Type:    protocol.MessageTypeError,
			Message: "config-lsp was unable to detect the appropriate language for this file. Please add: '#?lsp.language=<language>'.",
		},
	)
}

func NotifyDetectedLanguage(context *glsp.Context, uri protocol.DocumentUri, language shared.SupportedLanguage) {
	go context.Notify(
		"$/config-lsp/detectedLanguage",
		lspDetectedLanguage{
			lspNotification: lspNotification{
				Uri: string(uri),
			},
			Language: string(language),
		},
	)
}
