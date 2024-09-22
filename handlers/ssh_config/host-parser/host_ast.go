package hostparser

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
)

type Host struct {
	Hosts []*HostValue
}

type HostValue struct {
	common.LocationRange
	Value commonparser.ParsedString
}

