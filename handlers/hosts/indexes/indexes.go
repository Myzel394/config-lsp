package indexes

import "config-lsp/handlers/hosts/shared"

type HostsIndexes struct {
	Resolver *Resolver
	// [line]error
	DoubleIPs map[uint32]shared.DuplicateIPDeclaration
}

func NewHostsIndexes() HostsIndexes {
	return HostsIndexes{
		DoubleIPs: make(map[uint32]shared.DuplicateIPDeclaration),
	}
}
