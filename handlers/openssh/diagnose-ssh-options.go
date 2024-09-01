package openssh

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DiagnoseSSHOptions(
	context *glsp.Context,
	documentURI protocol.DocumentUri,
	parser *SimpleConfigParser,
) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	diagnostics = append(
		diagnostics,
		DiagnoseOption(
			context,
			documentURI,
			parser,
			"Port",
			func(value string, position SimpleConfigPosition) []protocol.Diagnostic {
				if value == "22" {
					severity := protocol.DiagnosticSeverityWarning

					return []protocol.Diagnostic{
						{
							Range: protocol.Range{
								Start: protocol.Position{
									Line:      position.Line,
									Character: uint32(len("Port ")),
								},
								End: protocol.Position{
									Line:      position.Line,
									Character: uint32(len("Port " + value)),
								},
							},
							Severity: &severity,
							Message:  "Port should not be 22 as it's often enumarated by attackers",
						},
					}
				}

				return []protocol.Diagnostic{}
			},
		)...,
	)

	return diagnostics
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
