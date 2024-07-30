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

func DiagnoseOption(
	context *glsp.Context,
	uri protocol.DocumentUri,
	parser *SimpleConfigParser,
	optionName string,
	checkerFunc func(string, SimpleConfigPosition) []protocol.Diagnostic,
) []protocol.Diagnostic {
	option, err := parser.GetOption(optionName)

	if err != nil {
		// Nothing to diagnose
		return nil
	}

	return checkerFunc(option.Value, option.Position)
}
