package ast

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/aliases/fields"
	"config-lsp/utils"
	"fmt"
)

type AliasValueInterface interface {
	GetAliasValue() AliasValue
}

func (a AliasValue) String() string {
	return fmt.Sprintf("%s %s", a.Location, a.Value)
}

func (a AliasValue) GetAliasValue() AliasValue {
	return a
}

type AliasValue struct {
	Location common.LocationRange
	Value    string
}

type AliasValueUser struct {
	AliasValue
}

type path string

type AliasValueFile struct {
	AliasValue
	Path path
}

func (a AliasValueFile) CheckIsValid() []common.LSPError {
	return utils.Map(
		fields.PathField.CheckIsValid(string(a.Path)),
		func(invalidValue *docvalues.InvalidValue) common.LSPError {
			return docvalues.LSPErrorFromInvalidValue(a.Location.Start.Line, *invalidValue)
		},
	)
}

type AliasValueCommand struct {
	AliasValue
	Command string
}

func (a AliasValueCommand) CheckIsValid() []common.LSPError {
	return utils.Map(
		fields.CommandField.CheckIsValid(a.Command),
		func(invalidValue *docvalues.InvalidValue) common.LSPError {
			return docvalues.LSPErrorFromInvalidValue(a.Location.Start.Line, *invalidValue)
		},
	)
}

type AliasValueIncludePath struct {
	Location common.LocationRange
	Path     path
}

type AliasValueInclude struct {
	AliasValue
	Path AliasValueIncludePath
}

func (a AliasValueInclude) CheckIsValid() []common.LSPError {
	return utils.Map(
		fields.PathField.CheckIsValid(string(a.Path.Path)),
		func(invalidValue *docvalues.InvalidValue) common.LSPError {
			return docvalues.LSPErrorFromInvalidValue(a.Location.Start.Line, *invalidValue)
		},
	)
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

type AliasValueError struct {
	AliasValue

	Code    *AliasValueErrorCode
	Message *AliasValueErrorMessage
}

type AliasValueErrorCode struct {
	AliasValue
}

type AliasValueErrorMessage struct {
	AliasValue
}
