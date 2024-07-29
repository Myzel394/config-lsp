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

	diagnostics := make([]protocol.Diagnostic, 0)

	diagnostics = append(
		diagnostics,
		common.Map(
			Parser.ParseFromFile(string(readBytes)),
			func (err common.OptionError) protocol.Diagnostic {
				return err.GetPublishDiagnosticsParams()
			},
		)...,
	)

	diagnostics = append(
		diagnostics,
		common.Map(
			common.AnalyzeValues(Parser, Options),
			func (err common.ValueError) protocol.Diagnostic {
				return err.GetPublishDiagnosticsParams()
			},
		)...,
	)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
