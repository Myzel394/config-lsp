package lsp

import (
	wireguard "config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormatting(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
	d := wireguard.DocumentParserMap[params.TextDocument.URI]

	return handlers.FormatDocument(
		d,
		params.Range,
		params.Options,
	)
}
