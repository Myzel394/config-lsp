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
	AvailableValues []string
	ProvidedValue   string
}

type NotANumberError struct{}

func (e NotANumberError) Error() string {
	return "This is not a number"
}

type NumberIsNotPositiveError struct{}

func (e NumberIsNotPositiveError) Error() string {
	return "This number is not positive"
}

type EmptyStringError struct{}

func (e EmptyStringError) Error() string {
	return "This setting may not be empty"
}

type ArrayContainsDuplicatesError struct {
	Duplicates []string
}

func (e ArrayContainsDuplicatesError) Error() string {
	return fmt.Sprintf("Array contains the following duplicate values: %s", strings.Join(e.Duplicates, ","))
}

type PathDoesNotExistError struct{}

func (e PathDoesNotExistError) Error() string {
	return "This path does not exist"
}

type PathInvalidError struct{}

func (e PathInvalidError) Error() string {
	return "This path is invalid"
}

type ValueError struct {
	Line  int
	Start int
	End   int
	Error error
}

func (e ValueNotInEnumError) Error() string {
	return fmt.Sprintf("'%s' is not valid. Select one from: %s", e.ProvidedValue, strings.Join(e.AvailableValues, ","))
}
