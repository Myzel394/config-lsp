package ast

import (
	"config-lsp/common"

	"github.com/emirpasic/gods/maps/treemap"
)

type SSHBlockType uint8

const (
	SSHBlockTypeMatch SSHBlockType = iota
	SSHBlockTypeHost
)

type SSHBlock interface {
	GetBlockType() SSHBlockType
	AddOption(option *SSHOption)
	SetEnd(common.Location)
	GetOptions() *treemap.Map
	GetEntryOption() *SSHOption
	GetLocation() common.LocationRange
}

func (b *SSHMatchBlock) GetBlockType() SSHBlockType {
	return SSHBlockTypeMatch
}

func (b *SSHMatchBlock) AddOption(option *SSHOption) {
	b.Options.Put(option.LocationRange.Start.Line, option)
}

func (b *SSHMatchBlock) SetEnd(end common.Location) {
	b.LocationRange.End = end
}

func (b *SSHMatchBlock) GetOptions() *treemap.Map {
	return b.Options
}

func (b *SSHMatchBlock) GetEntryOption() *SSHOption {
	return b.MatchOption
}

func (b *SSHMatchBlock) GetLocation() common.LocationRange {
	return b.LocationRange
}

func (b *SSHHostBlock) GetBlockType() SSHBlockType {
	return SSHBlockTypeHost
}

func (b *SSHHostBlock) AddOption(option *SSHOption) {
	b.Options.Put(option.LocationRange.Start.Line, option)
}

func (b *SSHHostBlock) SetEnd(end common.Location) {
	b.LocationRange.End = end
}

func (b *SSHHostBlock) GetOptions() *treemap.Map {
	return b.Options
}

func (b *SSHHostBlock) GetEntryOption() *SSHOption {
	return b.HostOption
}

func (b *SSHHostBlock) GetLocation() common.LocationRange {
	return b.LocationRange
}

type SSHType uint8

const (
	SSHTypeOption SSHType = iota
	SSHTypeMatch
	SSHTypeHost
)

type SSHEntry interface {
	GetType() SSHType
	GetOption() *SSHOption
}

func (o *SSHOption) GetType() SSHType {
	return SSHTypeOption
}

func (o *SSHOption) GetOption() *SSHOption {
	return o
}

func (b *SSHMatchBlock) GetType() SSHType {
	return SSHTypeMatch
}

func (b *SSHMatchBlock) GetOption() *SSHOption {
	return b.MatchOption
}

func (b *SSHHostBlock) GetType() SSHType {
	return SSHTypeHost
}

func (b *SSHHostBlock) GetOption() *SSHOption {
	return b.HostOption
}

func (c SSHConfig) FindBlock(line uint32) SSHBlock {
	it := c.Options.Iterator()
	for it.Next() {
		entry := it.Value().(SSHEntry)

		if entry.GetType() == SSHTypeOption {
			continue
		}

		block := entry.(SSHBlock)

		if block.GetLocation().Start.Line <= line && block.GetLocation().End.Line >= line {
			return block
		}
	}

	return nil
}

func (c SSHConfig) FindOption(line uint32) (*SSHOption, SSHBlock) {
	block := c.FindBlock(line)

	var option *SSHOption

	if block == nil {
		if rawOption, found := c.Options.Get(line); found {
			option = rawOption.(*SSHOption)
		}
	} else {
		if rawOption, found := block.GetOptions().Get(line); found {
			option = rawOption.(*SSHOption)
		}
	}

	return option, block
}

type AllOptionInfo struct {
	Block SSHBlock
	Option *SSHOption
}

func (c SSHConfig) GetAllOptions() []AllOptionInfo {
	options := make([]AllOptionInfo, 0, 50)

	for _, rawEntry := range c.Options.Values() {
		switch rawEntry.(type) {
		case *SSHOption:
			option := rawEntry.(*SSHOption)
			options = append(options, AllOptionInfo{
				Block: nil,
				Option: option,
			})
		case *SSHMatchBlock:
			block := rawEntry.(SSHBlock)

			options = append(options, AllOptionInfo{
				Block: block,
				Option: block.GetEntryOption(),
			})

			for _, rawOption := range block.GetOptions().Values() {
				option := rawOption.(*SSHOption)
				options = append(options, AllOptionInfo{
					Block: nil,
					Option: option,
				})
			}
		case *SSHHostBlock:
			block := rawEntry.(SSHBlock)

			options = append(options, AllOptionInfo{
				Block: block,
				Option: block.GetEntryOption(),
			})

			for _, rawOption := range block.GetOptions().Values() {
				option := rawOption.(*SSHOption)
				options = append(options, AllOptionInfo{
					Block: nil,
					Option: option,
				})
			}
		}

	}

	return options
}

