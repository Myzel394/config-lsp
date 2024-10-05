package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/ast"
	"fmt"
)

type HoverTarget string

const (
	HoverTargetIPAddress HoverTarget = "ip_address"
	HoverTargetHostname  HoverTarget = "hostname"
	HoverTargetAlias     HoverTarget = "alias"
)

func GetHoverTargetInEntry(
	index common.IndexPosition,
	e ast.HostsEntry,
) *HoverTarget {
	if e.IPAddress != nil && e.IPAddress.Location.ContainsPosition(index) {
		target := HoverTargetIPAddress
		return &target
	}

	if e.Hostname != nil && e.Hostname.Location.ContainsPosition(index) {
		target := HoverTargetHostname
		return &target
	}

	for _, alias := range e.Aliases {
		if alias.Location.ContainsPosition(index) {
			target := HoverTargetAlias
			return &target
		}
	}

	return nil
}

func GetHoverInfoForHostname(
	index common.IndexPosition,
	d hosts.HostsDocument,
	hostname ast.HostsHostname,
) []string {
	ipAddress := d.Indexes.Resolver.Entries[hostname.Value]

	return []string{
		fmt.Sprintf("**%s** maps to _%s_", hostname.Value, ipAddress.GetInfo()),
	}
}
