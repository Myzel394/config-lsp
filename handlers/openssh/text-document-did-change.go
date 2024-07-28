package handlers

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Todo: Implement incremental parsing
func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text

	Parser.Clear()
	errors := Parser.ParseFromFile(content)

	if len(errors) > 0 {
		SendDiagnosticsFromParserErrors(context, params.TextDocument.URI, errors)
	} else {
		ClearDiagnostics(context, params.TextDocument.URI)
	}


	return nil
}
