package roothandler

import (
	"config-lsp/handlers/fstab"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentDidChange(context, params)
	case LanguageSSHDConfig:
		return nil
	}

	panic("root-handler/TextDocumentDidChange: unexpected language" + language)
}
