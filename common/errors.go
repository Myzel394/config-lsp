package common

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type LSPError struct {
	Range LocationRange
	Err   error
}

func (l LSPError) ToDiagnostic() protocol.Diagnostic {
	return protocol.Diagnostic{
		Range:   l.Range.ToLSPRange(),
		Message: l.Err.Error(),
	}
}

type SyntaxError struct {
	Message string
}

func (s SyntaxError) Error() string {
	return s.Message
}
