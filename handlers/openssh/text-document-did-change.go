package handlers

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Todo: Implement incremental parsing
func TextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
		content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text

		Parser.Clear()
		Parser.ParseFromFile(content)

		return nil
}
