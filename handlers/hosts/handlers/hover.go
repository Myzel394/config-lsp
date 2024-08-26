package handlers

import (
	"config-lsp/handlers/hosts/handlers/analyzer"
	"fmt"
)

type HoverTarget string

const (
	HoverTargetIPAddress HoverTarget = "ip_address"
	HoverTargetHostname  HoverTarget = "hostname"
	HoverTargetAlias     HoverTarget = "alias"
)

func GetHoverTargetInEntry(
	e analyzer.HostsEntry,
	cursor uint32,
) *HoverTarget {
	if e.IPAddress != nil && e.IPAddress.Location.ContainsCursorByCharacter(cursor) {
		target := HoverTargetIPAddress
		return &target
	}

	if e.Hostname != nil && e.Hostname.Location.ContainsCursorByCharacter(cursor) {
		target := HoverTargetHostname
		return &target
	}

	for _, alias := range e.Aliases {
		if alias.Location.ContainsCursorByCharacter(cursor) {
			target := HoverTargetAlias
			return &target
		}
	}

	return nil
}

func GetHoverInfoForHostname(
	parser analyzer.HostsParser,
	hostname analyzer.HostsHostname,
	cursor uint32,
) []string {
	ipAddress := parser.Resolver.Entries[hostname.Value]

	return []string{
		fmt.Sprintf("**%s** maps to _%s_", hostname.Value, ipAddress.GetInfo()),
	}
}
