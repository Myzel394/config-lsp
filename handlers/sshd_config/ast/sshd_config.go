package ast

import (
	"config-lsp/common"
	match_parser "config-lsp/handlers/sshd_config/fields/match-parser"

	"github.com/emirpasic/gods/maps/treemap"
)

type SSHDKey struct {
	common.LocationRange
	Value string
}

type SSHDValue struct {
	common.LocationRange
	Value string
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
}

type SSHDOption struct {
	common.LocationRange
	Value string

	Key         *SSHDKey
	Separator   *SSHDSeparator
	OptionValue *SSHDValue
}

type SSHDMatchBlock struct {
	common.LocationRange
	MatchEntry *SSHDOption
	MatchValue *match_parser.Match

	// [uint32]*SSHDOption -> line number -> *SSHDOption
	Options *treemap.Map
}

type SSHDConfig struct {
	// [uint32]SSHDOption -> line number -> *SSHDEntry
	Options *treemap.Map
	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
