package utils

import (
	"context"
	"errors"
	net "net/netip"
)

type iPv4Tree struct {
	TrueNode  *iPv4Tree
	FalseNode *iPv4Tree
	Context   context.Context
}

func (t *iPv4Tree) addHostBits(
	hostBits []bool,
	ctx context.Context,
) {
	if len(hostBits) == 0 {
		t.Context = ctx
		return
	}

	if hostBits[0] {
		if t.TrueNode == nil {
			t.TrueNode = &iPv4Tree{}
		}
		t.TrueNode.addHostBits(hostBits[1:], ctx)
	} else {
		if t.FalseNode == nil {
			t.FalseNode = &iPv4Tree{}
		}
		t.FalseNode.addHostBits(hostBits[1:], ctx)
	}
}

func (t *iPv4Tree) getFromHostBits(hostBits []bool) *context.Context {
	if t.Context != nil || len(hostBits) == 0 {
		return &t.Context
	}

	if hostBits[0] {
		if t.TrueNode == nil {
			return nil
		}

		return t.TrueNode.getFromHostBits(hostBits[1:])
	} else {
		if t.FalseNode == nil {
			return nil
		}

		return t.FalseNode.getFromHostBits(hostBits[1:])
	}
}

func createIPv4Tree(
	hostBits []bool,
	ctx context.Context,
) iPv4Tree {
	tree := iPv4Tree{}
	tree.addHostBits(hostBits, ctx)

	return tree
}

type IPv4HostSet struct {
	tree iPv4Tree
}

func CreateIPv4HostSet() IPv4HostSet {
	return IPv4HostSet{
		tree: iPv4Tree{},
	}
}

// AddIP Add a new ip to the host set
// `hostAmount`: Amount of host bits
// Return: (<Whether the ip has been added>, <error>)
func (h *IPv4HostSet) AddIP(
	ip net.Prefix,
	ctx context.Context,
) (bool, error) {
	hostBits, err := ipToHostBits(ip)

	if err != nil {
		return false, err
	}

	if h.tree.getFromHostBits(hostBits) != nil {
		return false, nil
	}

	h.tree.addHostBits(hostBits, ctx)

	return true, nil
}

func (h IPv4HostSet) ContainsIP(
	ip net.Prefix,
) (*context.Context, error) {
	hostBits, err := ipToHostBits(ip)

	if err != nil {
		ctx := context.Background()
		return &ctx, err
	}

	ctx := h.tree.getFromHostBits(hostBits)

	return ctx, nil
}

func ipToHostBits(ip net.Prefix) ([]bool, error) {
	if !ip.Addr().Is4() {
		return nil, errors.New("Only IPv4 is supported currently")
	}

	ipv4 := ip.Addr().As4()
	allHostBits := [32]bool{}
	for i, b := range ipv4 {
		bits := byteToBits(b)
		for j, bit := range bits {
			allHostBits[i*8+j] = bit
		}
	}

	hostBits := allHostBits[:ip.Bits()]

	return hostBits, nil
}

func byteToBits(b byte) [8]bool {
	return [8]bool{
		(b>>0)&1 != 0,
		(b>>1)&1 != 0,
		(b>>2)&1 != 0,
		(b>>3)&1 != 0,
		(b>>4)&1 != 0,
		(b>>5)&1 != 0,
		(b>>6)&1 != 0,
		(b>>7)&1 != 0,
	}
}
