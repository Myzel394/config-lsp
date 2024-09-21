package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/common/parsers/openssh-match-parser"
	"github.com/emirpasic/gods/maps/treemap"
)

type SSHDKey struct {
	common.LocationRange
	Value commonparser.ParsedString
	Key   string
}

type SSHDValue struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type SSHDEntryType uint

const (
	SSHDEntryTypeOption SSHDEntryType = iota
	SSHDEntryTypeMatchBlock
)

type SSHDEntry interface {
	GetType() SSHDEntryType
	GetOption() SSHDOption
}

type SSHDSeparator struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type SSHDOption struct {
	common.LocationRange
	Value commonparser.ParsedString

	Key         *SSHDKey
	Separator   *SSHDSeparator
	OptionValue *SSHDValue
}

type SSHDMatchBlock struct {
	common.LocationRange
	MatchOption *SSHDOption
	MatchValue  *matchparser.Match

	// [uint32]*SSHDOption -> line number -> *SSHDOption
	Options *treemap.Map
}

type SSHDConfig struct {
	// [uint32]SSHDOption -> line number -> *SSHDEntry
	Options *treemap.Map
	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
