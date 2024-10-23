package utils

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type lspNotification struct {
	uri string
}

type lspDetectedLanguage struct {
	lspNotification

	language string
}

func NotifyLanguageUndetectable(context *glsp.Context, uri protocol.DocumentUri) {
	go context.Notify(
		"$/config-lsp/languageUndetectable",
		lspNotification{
			uri: string(uri),
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

func NotifyDetectedLanguage(context *glsp.Context, uri protocol.DocumentUri, language SupportedLanguage) {
	go context.Notify(
		"$/config-lsp/detectedLanguage",
		lspDetectedLanguage{
			lspNotification: lspNotification{
				uri: string(uri),
			},
			language: string(language),
		},
	)
}
