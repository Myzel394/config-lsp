package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type AllowedStatus uint8

const (
	AllowedStatusDisallowed AllowedStatus = iota
	AllowedStatusAllowed
	AllowedStatusRequired
)

type DeprecatedValue interface {
	GetTypeDescription() []string
	DeprecatedCheckIsValid(value string) []*InvalidValue
	DeprecatedFetchHoverInfo(line string, cursor uint32) []string

	FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem
}

type InvalidValue struct {
	Err   error
	Start uint32
	End   uint32
}

func (v *InvalidValue) Shift(offset uint32) {
	v.Start += offset
	v.End += offset
}

func (v *InvalidValue) GetRange(line uint32, characterStart uint32) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      line,
			Character: characterStart + v.Start,
		},
		End: protocol.Position{
			Line:      line,
			Character: characterStart + v.End + 1,
		},
	}
}

func (v *InvalidValue) GetMessage() string {
	return v.Err.Error()
}

func ShiftInvalidValues(offset uint32, invalidValues []*InvalidValue) {
	if len(invalidValues) > 0 {
		for _, invalidValue := range invalidValues {
			invalidValue.Shift(offset)
		}
	}
}

func InvalidValuesToErrorDiagnostics(
	line uint32,
	offset uint32,
	values []*InvalidValue,
) []protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError

	return utils.Map(
		values,
		func(value *InvalidValue) protocol.Diagnostic {
			return protocol.Diagnostic{
				Range:    value.GetRange(line, offset),
				Severity: &severity,
				Message:  value.GetMessage(),
			}
		},
	)
}
