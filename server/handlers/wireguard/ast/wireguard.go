package ast

import (
	"config-lsp/common"
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
	// map of: line number -> WGProperty
	Properties map[uint32]*WGProperty
}

type WGConfig struct {
	Sections []*WGSection
	// Used to identify where not to show diagnostics
	CommentLines map[uint32]struct{}
}
