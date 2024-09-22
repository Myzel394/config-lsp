package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	hostparser "config-lsp/handlers/ssh_config/host-parser"
	"config-lsp/handlers/ssh_config/match-parser"

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
	HostValue  *hostparser.Host

	// [uint32]*SSHOption -> line number -> *SSHOption
	Options *treemap.Map
}

type SSHConfig struct {
	// [uint32]SSHOption -> line number -> *SSHEntry
	Options *treemap.Map

	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
