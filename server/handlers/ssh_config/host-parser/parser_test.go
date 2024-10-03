package hostparser

import "testing"

func TestSimpleExample(
	t *testing.T,
) {
	input := `example.com`

	host := NewHost()
	offset := uint32(8)
	errs := host.Parse(input, 4, offset)

	if len(errs) > 0 {
		t.Fatalf("Expected no errors, got %v", errs)
	}

	if !(len(host.Hosts) == 1) {
		t.Errorf("Expected 1 host, got %v", len(host.Hosts))
	}

	if !(host.Hosts[0].Value.Raw == "example.com") {
		t.Errorf("Expected host to be 'example.com', got %v", host.Hosts[0].Value.Raw)
	}

	if !(host.Hosts[0].Start.Line == 4 && host.Hosts[0].Start.Character == 0+offset && host.Hosts[0].End.Character == 11+offset) {
		t.Errorf("Expected host to be at line 4, characters 0-11, got %v", host.Hosts[0])
	}

	if !(host.Hosts[0].Value.Value == "example.com") {
		t.Errorf("Expected host value to be 'example.com', got %v", host.Hosts[0].Value.Value)
	}
}

func TestMultipleExample(
	t *testing.T,
) {
	input := `example.com example.org example.net`

	host := NewHost()
	offset := uint32(8)
	errs := host.Parse(input, 4, offset)

	if len(errs) > 0 {
		t.Fatalf("Expected no errors, got %v", errs)
	}

	if !(len(host.Hosts) == 3) {
		t.Errorf("Expected 3 hosts, got %v", len(host.Hosts))
	}
}

func TestIncompleteExample(
	t *testing.T,
) {
	input := `example.com `

	host := NewHost()
	offset := uint32(8)
	errs := host.Parse(input, 4, offset)

	if len(errs) > 0 {
		t.Fatalf("Expected no errors, got %v", errs)
	}

	if !(len(host.Hosts) == 1) {
		t.Errorf("Expected 1 hosts, got %v", len(host.Hosts))
	}

	if !(host.Hosts[0].Value.Raw == "example.com") {
		t.Errorf("Expected host to be 'example.com', got %v", host.Hosts[0].Value.Raw)
	}
}
