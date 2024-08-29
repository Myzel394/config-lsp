package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/indexes"
	"config-lsp/handlers/hosts/shared"
	"config-lsp/utils"
	"net"
)

func createEntry(
	line uint32,
	ip net.IP,
) indexes.ResolverEntry {
	entry := indexes.ResolverEntry{
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

func createResolverFromParser(p ast.HostsParser) (indexes.Resolver, []common.LSPError) {
	errors := make([]common.LSPError, 0)
	resolver := indexes.Resolver{
		Entries: make(map[string]indexes.ResolverEntry),
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
					func(alias *ast.HostsHostname) hostnameEntry {
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
							Err: shared.DuplicateHostEntry{
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

func analyzeDoubleHostNames(d *hosts.HostsDocument) []common.LSPError {
	resolver, errors := createResolverFromParser(*d.Parser)

	d.Indexes.Resolver = &resolver

	return errors
}
