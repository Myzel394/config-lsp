package ast

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/aliases/fields"
	"config-lsp/utils"
	"fmt"
	"strconv"
)

type AliasValueInterface interface {
	GetAliasValue() AliasValue
	GetStructName() string
}

func (a AliasValue) String() string {
	return fmt.Sprintf("%s %s", a.Location, a.Value)
}

func (a AliasValue) GetAliasValue() AliasValue {
	return a
}

func (a AliasValue) GetStructName() string {
	return "AliasValue"
}

type AliasValue struct {
	Location common.LocationRange
	Value    string
}

type AliasValueUser struct {
	AliasValue
}

func (a AliasValueUser) GetStructName() string {
	return "AliasValueUser"
}

type path string

type AliasValueFile struct {
	AliasValue
	Path path
}

func (a AliasValueFile) GetStructName() string {
	return "AliasValueFile"
}

type AliasValueCommand struct {
	AliasValue
	Command string
}

func (a AliasValueCommand) GetStructName() string {
	return "AliasValueCommand"
}

type AliasValueIncludePath struct {
	Location common.LocationRange
	Path     path
}

type AliasValueInclude struct {
	AliasValue
	Path *AliasValueIncludePath
}

func (a AliasValueInclude) CheckIsValid() []common.LSPError {
	return utils.Map(
		fields.PathField.CheckIsValid(string(a.Path.Path)),
		func(invalidValue *docvalues.InvalidValue) common.LSPError {
			return docvalues.LSPErrorFromInvalidValue(a.Location.Start.Line, *invalidValue)
		},
	)
}

func (a AliasValueInclude) GetStructName() string {
	return "AliasValueInclude"
}

type AliasValueEmail struct {
	AliasValue
}

func (a AliasValueEmail) CheckIsValid() []common.LSPError {
	return utils.Map(
		fields.PathField.CheckIsValid(a.Value),
		func(invalidValue *docvalues.InvalidValue) common.LSPError {
			return docvalues.LSPErrorFromInvalidValue(a.Location.Start.Line, *invalidValue)
		},
	)
}

func (a AliasValueEmail) GetStructName() string {
	return "AliasValueEmail"
}

type AliasValueError struct {
	AliasValue

	Code    *AliasValueErrorCode
	Message *AliasValueErrorMessage
}

type AliasValueErrorCode struct {
	AliasValue
}

func (a AliasValueError) GetStructName() string {
	return "AliasValueError"
}

func (a AliasValueErrorCode) ErrorCodeAsInt() uint16 {
	code, err := strconv.Atoi(a.Value)

	if err != nil {
		return 0
	}

	return uint16(code)
}

type AliasValueErrorMessage struct {
	AliasValue
}
