package roothandler

import (
	"config-lsp/handlers/fstab"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	language := rootHandler.GetLanguageForDocument(params.TextDocument.URI)

	switch language {
	case LanguageFstab:
		return fstab.TextDocumentHover(context, params)
	case LanguageSSHDConfig:
		return nil, nil
	}

	panic("root-handler/TextDocumentHover: unexpected language" + language)
}
