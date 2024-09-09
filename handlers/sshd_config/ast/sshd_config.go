package ast

import (
	"config-lsp/common"

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
}

type SSHOption struct {
	common.LocationRange
	Value string

	Key         *SSHKey
	OptionValue *SSHValue
}

func (o SSHOption) GetType() SSHEntryType {
	return SSHEntryTypeOption
}

type SSHMatchBlock struct {
	common.LocationRange
	MatchEntry *SSHOption

	// [uint32]*SSHOption -> line number -> *SSHOption
	Options *treemap.Map
}

func (m SSHMatchBlock) GetType() SSHEntryType {
	return SSHEntryTypeMatchBlock
}

type SSHConfig struct {
	// [uint32]SSHOption -> line number -> *SSHEntry
	Options *treemap.Map
	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}
