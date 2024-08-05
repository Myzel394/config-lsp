package roothandler

import (
	"config-lsp/handlers/fstab"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentCompletion(context, params)
	case LanguageSSHDConfig:
		return nil, nil
	}

	panic("root-handler/TextDocumentCompletion: unexpected language" + language)
}
