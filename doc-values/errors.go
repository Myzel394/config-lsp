package docvalues

import (
	"config-lsp/common"
	"fmt"
	"unicode/utf8"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DiagnosticableError interface {
	GetDiagnostic(uri protocol.DocumentUri) protocol.Diagnostic
}

type OptionError struct {
	Line           uint32
	ProvidedOption string
	DocError       error
}

func (e OptionError) GetPublishDiagnosticsParams() protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError

	return protocol.Diagnostic{
		Message: e.DocError.Error(),
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      e.Line,
				Character: 0,
			},
			End: protocol.Position{
				Line:      e.Line,
				Character: uint32(utf8.RuneCountInString(e.ProvidedOption)),
			},
		},
		Severity: &severity,
	}
}
func (e OptionError) Error() string {
	return "Option error"
}

type ValueError struct {
	Line   uint32
	Option string
	Value  string

	DocError error
}

func (e ValueError) GetPublishDiagnosticsParams() protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError
	start := uint32(utf8.RuneCountInString(e.Option) + utf8.RuneCountInString(" "))

	return protocol.Diagnostic{
		Message: e.DocError.Error(),
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      e.Line,
				Character: start,
			},
			End: protocol.Position{
				Line:      e.Line,
				Character: start + uint32(utf8.RuneCountInString(e.Value)),
			},
		},
		Severity: &severity,
	}
}
func (e ValueError) Error() string {
	return "Value error"
}

type OptionAlreadyExistsError struct {
	AlreadyLine uint32
}

func (e OptionAlreadyExistsError) Error() string {
	return fmt.Sprintf("This option is already defined on line %d", e.AlreadyLine)
}

type OptionUnknownError struct{}

func (e OptionUnknownError) Error() string {
	return "This option does not exist"
}

type MalformedLineError struct{}

func (e MalformedLineError) Error() string {
	return "Malformed line"
}

type LineNotFoundError struct{}

func (e LineNotFoundError) Error() string {
	return "Line not found"
}

func LSPErrorFromInvalidValue(
	line uint32,
	invaludValue InvalidValue,
) common.LSPError {
	return common.LSPError{
		Range: common.LocationRange{
			Start: common.Location{
				Line:      line,
				Character: invaludValue.Start,
			},
			End: common.Location{
				Line:      line,
				Character: invaludValue.End,
			},
		},
		Err: invaludValue.Err,
	}
}
