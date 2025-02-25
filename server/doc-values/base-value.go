package docvalues

import (
	"config-lsp/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DeprecatedValue interface {
	GetTypeDescription() []string
	DeprecatedCheckIsValid(value string) []*InvalidValue
	DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem
	DeprecatedFetchHoverInfo(line string, cursor uint32) []string
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
