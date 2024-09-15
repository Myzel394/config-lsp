package ast

import (
	"config-lsp/common"
	match_parser "config-lsp/handlers/sshd_config/fields/match-parser"

	"github.com/emirpasic/gods/maps/treemap"
)

type SSHKey struct {
	common.LocationRange
	Value string
}

type SSHValue struct {
	common.LocationRange
	Value string
}

type SSHEntryType uint

const (
	SSHEntryTypeOption SSHEntryType = iota
	SSHEntryTypeMatchBlock
)

type SSHEntry interface {
	GetType() SSHEntryType
	GetOption() SSHOption
}

type SSHSeparator struct {
	common.LocationRange
}

type SSHOption struct {
	common.LocationRange
	Value string

	Key         *SSHKey
	Separator   *SSHSeparator
	OptionValue *SSHValue
}

type SSHMatchBlock struct {
	common.LocationRange
	MatchEntry *SSHOption
	MatchValue *match_parser.Match

	// [uint32]*SSHOption -> line number -> *SSHOption
	Options *treemap.Map
}

type SSHConfig struct {
	// [uint32]SSHOption -> line number -> *SSHEntry
	Options *treemap.Map
	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
