package ast

import (
	"config-lsp/common"
	"config-lsp/common/parser"

	"github.com/emirpasic/gods/maps/treemap"
)

type GitKey struct {
	common.LocationRange
	Value parser.ParsedString
}

type GitSeparator struct {
	common.LocationRange
	Value parser.ParsedString
}

type GitValuePart struct {
	common.LocationRange
	Value parser.ParsedString
}

type GitValue struct {
	Raw   common.VirtualLine
	Value string
}

type GitEntry struct {
	common.LocationRange
	Key       *GitKey
	Separator *GitSeparator
	Value     *GitValue
}

type GitSectionHeader struct {
	common.LocationRange
	Title string
}

type GitSection struct {
	common.LocationRange
	Entries *treemap.Map
	Title   *GitSectionHeader
}

type GitConfig struct {
	Sections     []*GitSection
	CommentLines map[uint32]struct{}
}
