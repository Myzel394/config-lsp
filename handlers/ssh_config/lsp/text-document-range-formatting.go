package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormatting(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
	return nil, nil
}
