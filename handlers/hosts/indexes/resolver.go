package indexes

import (
	"net"
)

type ResolverEntry struct {
	IPv4Address *net.IP
	IPv6Address *net.IP
	Line        uint32
}

func (e ResolverEntry) GetInfo() string {
	if e.IPv4Address != nil {
		return e.IPv4Address.String()
	}

	return e.IPv6Address.String()
}

type Resolver struct {
	Entries map[string]*ResolverEntry
}
