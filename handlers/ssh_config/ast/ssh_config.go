package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/common/parsers/openssh-match-parser"
	"github.com/emirpasic/gods/maps/treemap"
)

type SSHKey struct {
	common.LocationRange
	Value commonparser.ParsedString
	Key   string
}

type SSHSeparator struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type SSHValue struct {
	common.LocationRange
	Value commonparser.ParsedString
}

type SSHOption struct {
	common.LocationRange
	Value commonparser.ParsedString

	Key         *SSHKey
	Separator   *SSHSeparator
	OptionValue *SSHValue
}

type SSHMatchBlock struct {
	common.LocationRange
	MatchOption *SSHOption
	MatchValue  *matchparser.Match

	// [uint32]*SSHOption -> line number -> *SSHOption
	Options *treemap.Map
}

type SSHHostBlock struct {
	common.LocationRange
	HostOption *SSHOption
	HostValue  string

	// [uint32]*SSHOption -> line number -> *SSHOption
	Others *treemap.Map
}

type SSHConfig struct {
	// [uint32]SSHOption -> line number -> *SSHEntry
	RootOptions *treemap.Map

	MatchBlosks []*SSHMatchBlock
	HostBlocks  []*SSHHostBlock

	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
