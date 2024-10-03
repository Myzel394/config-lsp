package common

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ParseError struct {
	Line uint32
	Err  error
}

func (e ParseError) Error() string {
	return "Parse error"
}
func (e ParseError) ToDiagnostic() protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError
	return protocol.Diagnostic{
		Severity: &severity,
		Message:  e.Err.Error(),
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      e.Line,
				Character: 0,
			},
			End: protocol.Position{
				Line:      e.Line,
				Character: 999999,
			},
		},
	}
}
