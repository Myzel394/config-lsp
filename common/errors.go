package common

import (
	"fmt"
	"strings"
)

type ParserError interface{}

type OptionAlreadyExistsError struct {
	Option      string
	FoundOnLine uint32
}

func (e OptionAlreadyExistsError) Error() string {
	return fmt.Sprintf("Option %s already exists", e.Option)
}

type OptionUnknownError struct {
	Option string
}

func (e OptionUnknownError) Error() string {
	return fmt.Sprintf("Option '%s' does not exist", e.Option)
}

type MalformedLineError struct {
	Line string
}

func (e MalformedLineError) Error() string {
	return fmt.Sprintf("Malformed line: %s", e.Line)
}

type LineNotFoundError struct{}

func (e LineNotFoundError) Error() string {
	return "Line not found"
}

type ValueNotInEnumError struct {
	availableValues []string
	providedValue   string
}

func (e ValueNotInEnumError) Error() string {
	return fmt.Sprint("'%s' is not valid. Select one from: %s", e.providedValue, strings.Join(e.availableValues, ","))
}
