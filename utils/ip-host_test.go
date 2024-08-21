package utils

import (
	"context"
	"net/netip"
	"testing"
)

func TestFullHostIpAddresses(t *testing.T) {
	// Test the full host IP address
	hostSet := CreateIPv4HostSet()

	hostSet.AddIP(netip.MustParsePrefix("10.0.0.1/32"), context.Background())
	hostSet.AddIP(netip.MustParsePrefix("10.0.0.2/32"), context.Background())
	hostSet.AddIP(netip.MustParsePrefix("10.0.0.3/32"), context.Background())

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("10.0.0.1/32")); ctx == nil {
		t.Fatalf("Expected to find 10.0.0.1/32 in the host set")
	}

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("10.0.0.5/32")); ctx != nil {
		t.Fatalf("Expected NOT to find 10.0.0.5/32 in the host set")
	}
}

func TestPartialHostIpAddresses(t *testing.T) {
	// Test the partial host IP address
	hostSet := CreateIPv4HostSet()

	hostSet.AddIP(netip.MustParsePrefix("10.0.0.1/32"), context.Background())
	hostSet.AddIP(netip.MustParsePrefix("10.0.0.2/32"), context.Background())
	hostSet.AddIP(netip.MustParsePrefix("10.0.0.3/32"), context.Background())

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("10.0.0.1/16")); ctx == nil {
		t.Fatalf("Expected to find 10.0.0.1/16 in the host set")
	}

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("192.168.0.1/16")); ctx != nil {
		t.Fatalf("Expected NOT to find 192.168.0.1/16 in the host set")
	}
}

func TestMixedHostIpAddresses(t *testing.T) {
	// Test the mixed host IP address
	hostSet := CreateIPv4HostSet()

	hostSet.AddIP(netip.MustParsePrefix("10.0.0.1/16"), context.Background())
	hostSet.AddIP(netip.MustParsePrefix("192.168.0.1/32"), context.Background())

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("10.0.0.2/32")); ctx == nil {
		t.Fatalf("Expected to find 10.0.0.3/32 in the host set")
	}

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("192.168.0.2/32")); ctx != nil {
		t.Fatalf("Expected NOT to find 192.168.0.2/32 in the host set")
	}
}

func TestSimpleExactCheck(t *testing.T) {
	// Test the real example
	hostSet := CreateIPv4HostSet()

	hostSet.AddIP(netip.MustParsePrefix("10.0.0.1/16"), context.Background())

	if ctx, _ := hostSet.ContainsIP(netip.MustParsePrefix("10.0.0.1/16")); ctx == nil {
		t.Fatalf("Expected to find 10.0.0.1/16 in the host set")
	}
}
