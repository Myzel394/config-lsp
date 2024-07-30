package openssh

import (
	"config-lsp/common"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DiagnoseSSHOptions(
	context *glsp.Context,
	documentURI protocol.DocumentUri,
	parser *common.SimpleConfigParser,
) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	diagnostics = append(
		diagnostics,
		common.DiagnoseOption(
			context,
			documentURI,
			parser,
			"Port",
			func(value string, position common.SimpleConfigPosition) []protocol.Diagnostic {
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
