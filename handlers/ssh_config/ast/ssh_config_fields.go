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
