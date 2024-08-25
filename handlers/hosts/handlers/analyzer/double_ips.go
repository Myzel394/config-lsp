package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts/tree"
	"net"
)

func ipToString(ip net.IPAddr) string {
	return ip.IP.String()
}

func analyzeDoubleIPs(p tree.HostsParser) []common.LSPError {
	errors := make([]common.LSPError, 0)

	ips := make(map[string]uint32)

	for lineNumber, entry := range p.Tree.Entries {
		if entry.IPAddress != nil {
			key := ipToString(entry.IPAddress.Value)

			if foundLine, found := ips[key]; found {
				errors = append(errors, common.LSPError{
					Range: entry.IPAddress.Location,
					Err: DuplicateIPDeclaration{
						AlreadyFoundAt: foundLine,
					},
				})
			} else {
				ips[key] = lineNumber
			}
		}
	}

	return errors
}
