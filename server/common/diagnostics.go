package common

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func ClearDiagnostics(context *glsp.Context, uri protocol.DocumentUri) {
	context.Notify(
		"textDocument/publishDiagnostics",
		protocol.PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: []protocol.Diagnostic{},
		},
	)
}

func SendDiagnostics(context *glsp.Context, uri protocol.DocumentUri, diagnostics []protocol.Diagnostic) {
	context.Notify(
		"textDocument/publishDiagnostics",
		protocol.PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: diagnostics,
		},
	)
}
