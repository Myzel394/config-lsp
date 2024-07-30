package openssh

import (
	"config-lsp/common"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Todo: Implement incremental parsing
func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text

	Parser.Clear()
	diagnostics := DiagnoseParser(context, params.TextDocument.URI, content)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	} else {
		common.ClearDiagnostics(context, params.TextDocument.URI)
	}

	return nil
}
