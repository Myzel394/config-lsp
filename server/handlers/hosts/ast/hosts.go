package ast

import (
	"config-lsp/common"
	"fmt"
	"net"

	"github.com/emirpasic/gods/maps/treemap"
)

type HostsParser struct {
	Tree         HostsTree
	CommentLines map[uint32]struct{}
}

type HostsTree struct {
	// [line]entry
	Entries *treemap.Map
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
		str += " " + p.IPAddress.Value.String()
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
	Value    net.IPAddr
}

type HostsHostname struct {
	Location common.LocationRange
	Value    string
}
