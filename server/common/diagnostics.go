package common

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func ClearDiagnostics(context *glsp.Context, uri protocol.DocumentUri) {
	// Diagnostics are sent synchronously, as sending them async
	// could result in a race condition when we send diagnostics
	// to the client.
	context.Notify(
		protocol.ServerTextDocumentPublishDiagnostics,
		protocol.PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: []protocol.Diagnostic{},
		},
	)
}

func SendDiagnostics(context *glsp.Context, uri protocol.DocumentUri, diagnostics []protocol.Diagnostic) {
	go context.Notify(
		protocol.ServerTextDocumentPublishDiagnostics,
		protocol.PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: diagnostics,
		},
	)
}
