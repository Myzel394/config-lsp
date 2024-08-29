package tree

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type AliasValueInterface interface {
	FetchCompletions(line string, cursor uint32) []protocol.CompletionItem
	CheckIsValid() []*docvalues.InvalidValue
}

func (a AliasValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return nil
}
func (a AliasValue) CheckIsValid() []*docvalues.InvalidValue {
	return nil
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

type AliasValueCommand struct {
	AliasValue
	Command string
}

type AliasValueIncludePath struct {
	Location common.LocationRange
	Path     path
}

type AliasValueInclude struct {
	AliasValue
	Path AliasValueIncludePath
}

type AliasValueEmail struct {
	AliasValue
}
