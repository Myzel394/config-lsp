package openssh

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DiagnoseParser(
	context *glsp.Context,
	documentURI protocol.DocumentUri,
	content string,
) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	diagnostics = append(
		diagnostics,
		utils.Map(
			Parser.ParseFromFile(content),
			func(err docvalues.OptionError) protocol.Diagnostic {
				return err.GetPublishDiagnosticsParams()
			},
		)...,
	)

	diagnostics = append(
		diagnostics,
		utils.Map(
			AnalyzeValues(Parser, Options),
			func(err docvalues.ValueError) protocol.Diagnostic {
				return err.GetPublishDiagnosticsParams()
			},
		)...,
	)

	diagnostics = append(
		diagnostics,
		DiagnoseSSHOptions(
			context,
			documentURI,
			&Parser,
		)...,
	)

	return diagnostics
}
