package analyzer

import (
	"config-lsp/common"
	"net"
)

func ipToString(ip net.IPAddr) string {
	return ip.IP.String()
}

func analyzeDoubleIPs(p *HostsParser) []common.LSPError {
	errors := make([]common.LSPError, 0)
	ips := make(map[string]uint32)

	p.DoubleIPs = make(map[uint32]DuplicateIPDeclaration)

	for lineNumber, entry := range p.Tree.Entries {
		if entry.IPAddress != nil {
			key := ipToString(entry.IPAddress.Value)

			if foundLine, found := ips[key]; found {
				err := DuplicateIPDeclaration{
					AlreadyFoundAt: foundLine,
				}

				p.DoubleIPs[lineNumber] = err
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
