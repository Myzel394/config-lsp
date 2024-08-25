package tree

import (
	"config-lsp/common"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func characterRangeFromCtx(
	ctx antlr.BaseParserRuleContext,
) common.LocationRange {
	line := uint32(ctx.GetStart().GetLine())
	start := uint32(ctx.GetStart().GetStart())
	end := uint32(ctx.GetStop().GetStop())

	return common.LocationRange{
		Start: common.Location{
			Line:      line,
			Character: start,
		},
		End: common.Location{
			Line:      line,
			Character: end,
		},
	}
}

type HostsParser struct {
	Tree         HostsTree
	CommentLines map[uint32]struct{}
}

type HostsTree struct {
	// [line]entry
	Entries map[uint32]*HostsEntry
}

type HostsEntry struct {
	Location common.LocationRange

	IPAddress *HostsIPAddress
	Hostname  *HostsHostname
	Aliases   []*HostsHostname
}

func (p HostsEntry) String() string {
	str := fmt.Sprintf("HostsEntry(%v)", p.Location)

	if p.IPAddress != nil {
		str += " " + p.IPAddress.Value
	}

	if p.Hostname != nil {
		str += " " + p.Hostname.Value
	}

	if p.Aliases != nil {
		str += " " + fmt.Sprintf("%v", p.Aliases)
	}

	return str
}

type HostsIPAddress struct {
	Location common.LocationRange
	Value    string
}

type HostsHostname struct {
	Location common.LocationRange
	Value    string
}
