package ast

import (
	"config-lsp/common"
	"github.com/emirpasic/gods/maps/treemap"
)

type WGPropertyKey struct {
	common.LocationRange
	Name string
}

type WGPropertyValue struct {
	common.LocationRange
	Value string
}

type WGPropertySeparator struct {
	common.LocationRange
}

type WGProperty struct {
	common.LocationRange
	RawValue  string
	Key       WGPropertyKey
	Separator *WGPropertySeparator
	Value     *WGPropertyValue
}

type WGHeader struct {
	common.LocationRange
	Name string
}

type WGSection struct {
	common.LocationRange
	Header WGHeader
	// [uint32]*WGProperty: line number -> *WGProperty
	Properties *treemap.Map
}

type WGConfig struct {
	Sections []*WGSection
	// Used to identify where not to show diagnostics
	CommentLines map[uint32]struct{}
}
