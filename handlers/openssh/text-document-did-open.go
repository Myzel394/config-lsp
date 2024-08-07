package openssh

import (
	"config-lsp/common"
	"os"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	readBytes, err := os.ReadFile(params.TextDocument.URI[len("file://"):])

	if err != nil {
		return err
	}

	diagnostics := DiagnoseParser(context, params.TextDocument.URI, string(readBytes))

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
