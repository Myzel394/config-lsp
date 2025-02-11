package utils

import (
	"config-lsp/common"
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
	// We always send this undetectable message, as it's a custom message.
	// The client can handle it themselves.
	go context.Notify(
		"$/config-lsp/languageUndetectable",
		lspNotification{
			Uri: string(uri),
		},
	)

	// The native showMessage notification however, should only be shown
	// if the user wishes to.
	if !common.ServerOptions.NoUndetectableErrors {
		go context.Notify(
			"window/showMessage",
			protocol.ShowMessageParams{
				Type:    protocol.MessageTypeError,
				Message: "config-lsp was unable to detect the appropriate language for this file. Please add: '#?lsp.language=<language>'.",
			},
		)
	}
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
