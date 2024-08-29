package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/shared"
	"net"
)

func ipToString(ip net.IPAddr) string {
	return ip.IP.String()
}

func analyzeDoubleIPs(d *hosts.HostsDocument) []common.LSPError {
	errors := make([]common.LSPError, 0)
	ips := make(map[string]uint32)

	d.Indexes.DoubleIPs = make(map[uint32]shared.DuplicateIPDeclaration)

	for lineNumber, entry := range d.Parser.Tree.Entries {
		if entry.IPAddress != nil {
			key := ipToString(entry.IPAddress.Value)

			if foundLine, found := ips[key]; found {
				err := shared.DuplicateIPDeclaration{
					AlreadyFoundAt: foundLine,
				}

				d.Indexes.DoubleIPs[lineNumber] = err
				errors = append(errors, common.LSPError{
					Range: entry.IPAddress.Location,
					Err:   err,
				})
			} else {
				ips[key] = lineNumber
			}
		}
	}

	return errors
}
