package common

import (
	"fmt"
	"strings"
	"unicode/utf8"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DiagnosticableError interface {
	GetDiagnostic(uri protocol.DocumentUri) protocol.Diagnostic
}

type OptionError struct {
	Line uint32
	ProvidedOption string
	DocError error
}
func (e OptionError) GetPublishDiagnosticsParams() protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError

	return protocol.Diagnostic{
		Message: e.DocError.Error(),
		Range: protocol.Range{
			Start: protocol.Position{
				Line: e.Line,
				Character: 0,
			},
			End: protocol.Position{
				Line: e.Line,
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
	Line uint32
	Option string
	Value string

	DocError error
}
func (e ValueError) GetPublishDiagnosticsParams() protocol.Diagnostic {
	severity := protocol.DiagnosticSeverityError
	start := uint32(utf8.RuneCountInString(e.Option) + utf8.RuneCountInString(" "))

	return protocol.Diagnostic{
		Message: e.DocError.Error(),
		Range: protocol.Range{
			Start: protocol.Position{
				Line: e.Line,
				Character: start,
			},
			End: protocol.Position{
				Line: e.Line,
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

type OptionUnknownError struct {}

func (e OptionUnknownError) Error() string {
	return "This option does not exist"
}

type MalformedLineError struct {}

func (e MalformedLineError) Error() string {
	return "Malformed line"
}

type LineNotFoundError struct{}

func (e LineNotFoundError) Error() string {
	return "Line not found"
}

// Value errors
type ValueNotInEnumError struct {
	AvailableValues []string
	ProvidedValue   string
}

func (e ValueNotInEnumError) Error() string {
	return fmt.Sprintf("This value is not valid. Select one from: %s", strings.Join(e.AvailableValues, ","))
}

type NotANumberError struct{}

func (e NotANumberError) Error() string {
	return "This must be number"
}

type NumberIsNotPositiveError struct{}

func (e NumberIsNotPositiveError) Error() string {
	return "This number must be positive for this setting"
}

type EmptyStringError struct{}

func (e EmptyStringError) Error() string {
	return "This setting may not be empty"
}

type ArrayContainsDuplicatesError struct {
	Duplicates []string
}
func (e ArrayContainsDuplicatesError) Error() string {
	return fmt.Sprintf("Remove the following duplicate values: %s", strings.Join(e.Duplicates, ","))
}

type PathDoesNotExistError struct{}

func (e PathDoesNotExistError) Error() string {
	return "This path does not exist"
}

type KeyValueAssignmentError struct{}

func (e KeyValueAssignmentError) Error() string {
	return "This is not valid key-value assignment"
}

type PathInvalidError struct{}

func (e PathInvalidError) Error() string {
	return "This path is invalid"
}

