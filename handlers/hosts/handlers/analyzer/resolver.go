package analyzer

import (
	"config-lsp/common"
	"config-lsp/utils"
	"net"
)

type ResolverEntry struct {
	IPv4Address net.IP
	IPv6Address net.IP
	Line        uint32
}

func (e ResolverEntry) GetInfo() string {
	if e.IPv4Address != nil {
		return e.IPv4Address.String()
	}

	return e.IPv6Address.String()
}

type Resolver struct {
	Entries map[string]ResolverEntry
}

func createEntry(
	line uint32,
	ip net.IP,
) ResolverEntry {
	entry := ResolverEntry{
		Line: line,
	}

	if ipv4 := ip.To4(); ipv4 != nil {
		entry.IPv4Address = ipv4
	} else if ipv6 := ip.To16(); ipv6 != nil {
		entry.IPv6Address = ipv6
	}

	return entry
}

type hostnameEntry struct {
	Location common.LocationRange
	HostName string
}

func createResolverFromParser(p HostsParser) (Resolver, []common.LSPError) {
	errors := make([]common.LSPError, 0)
	resolver := Resolver{
		Entries: make(map[string]ResolverEntry),
	}

	for lineNumber, entry := range p.Tree.Entries {
		if entry.IPAddress != nil && entry.Hostname != nil {
			hostNames := append(
				[]hostnameEntry{
					{
						Location: entry.Hostname.Location,
						HostName: entry.Hostname.Value,
					},
				},
				utils.Map(
					entry.Aliases,
					func(alias *HostsHostname) hostnameEntry {
						return hostnameEntry{
							Location: alias.Location,
							HostName: alias.Value,
						}
					},
				)...,
			)

			for _, hostName := range hostNames {
				entry := createEntry(
					lineNumber,
					entry.IPAddress.Value.IP,
				)

				if resolv, found := resolver.Entries[hostName.HostName]; found {
					errors = append(
						errors,
						common.LSPError{
							Range: hostName.Location,
							Err: DuplicateHostEntry{
								AlreadyFoundAt: resolv.Line,
								Hostname:       hostName.HostName,
							},
						},
					)
				} else {
					resolver.Entries[hostName.HostName] = entry
				}
			}
		}
	}

	return resolver, errors
}

func analyzeDoubleHostNames(p *HostsParser) []common.LSPError {
	resolver, errors := createResolverFromParser(*p)

	p.Resolver = &resolver

	return errors
}
