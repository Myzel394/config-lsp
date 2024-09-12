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

func (o SSHOption) GetType() SSHEntryType {
	return SSHEntryTypeOption
}

func (o SSHOption) GetOption() SSHOption {
	return o
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

func (m SSHMatchBlock) GetOption() SSHOption {
	return *m.MatchEntry
}

type SSHConfig struct {
	// [uint32]SSHOption -> line number -> *SSHEntry
	Options *treemap.Map
	// [uint32]{} -> line number -> {}
	CommentLines map[uint32]struct{}
}

func (c SSHConfig) FindMatchBlock(line uint32) *SSHMatchBlock {
	for currentLine := line; currentLine > 0; currentLine-- {
		rawEntry, found := c.Options.Get(currentLine)

		if !found {
			continue
		}

		switch entry := rawEntry.(type) {
		case *SSHMatchBlock:
			return entry
		}
	}

	return nil
}

func (c SSHConfig) FindOption(line uint32) (*SSHOption, *SSHMatchBlock) {
	matchBlock := c.FindMatchBlock(line)

	if matchBlock != nil {
		rawEntry, found := matchBlock.Options.Get(line)

		if found {
			return rawEntry.(*SSHOption), matchBlock
		}
	}

	rawEntry, found := c.Options.Get(line)

	if found {
		switch rawEntry.(type) {
		case *SSHMatchBlock:
			return rawEntry.(*SSHMatchBlock).MatchEntry, rawEntry.(*SSHMatchBlock)
		case *SSHOption:
			return rawEntry.(*SSHOption), nil
		}
	}

	return nil, nil

}
