package lsp

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormatting(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
	d := sshconfig.DocumentParserMap[params.TextDocument.URI]

	return handlers.FormatDocument(
		d,
		params.Range,
		params.Options,
	)
}
