package handlers

import (
	"net"
	"testing"
)

func TestCreateNewIPSimple24Mask(t *testing.T) {
	_, network, _ := net.ParseCIDR("10.0.0.0/24")
	newIP := createNewIP(*network, "10.0.0.1/32")

	if newIP != "10.0.0.2/32" {
		t.Errorf("Expected 10.0.0.2/32, got %s", newIP)
	}
}

func TestCreateNewIPDoesNotWorkWithLast24Mask(t *testing.T) {
	_, network, _ := net.ParseCIDR("10.0.0.0/24")
	newIP := createNewIP(*network, "10.0.0.254/32")

	if newIP != "" {
		t.Errorf("Expected empty string, got %s", newIP)
	}
}

func TestCreateNewIPDoesNotWorkWithLast24Mask2(t *testing.T) {
	_, network, _ := net.ParseCIDR("10.0.0.0/24")
	newIP := createNewIP(*network, "10.0.0.255/32")

	if newIP != "" {
		t.Errorf("Expected empty string, got %s", newIP)
	}
}

func TestCreateNewIPComplex20Mask(t *testing.T) {
	_, network, _ := net.ParseCIDR("10.0.0.0/20")
	newIP := createNewIP(*network, "10.0.0.255/32")

	if newIP != "10.0.1.0/32" {
		t.Errorf("Expected 10.0.1.0/32, got %s", newIP)
	}
}
