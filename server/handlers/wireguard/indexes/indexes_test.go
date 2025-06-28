package indexes

import (
	"config-lsp/handlers/wireguard/ast"
	"testing"
)

func TestIndexesUpProperties(t *testing.T) {
	content := `
[Interface]
PrivateKey = abc
PreUp = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT
`

	config := ast.NewWGConfig()
	error := config.Parse(content)

	if len(error) > 0 {
		t.Fatalf("Failed to parse config: %v", error)
	}

	indexes, errs := CreateIndexes(config)

	if len(errs) > 0 {
		t.Fatalf("Unexpected errors while creating indexes: %v", errs)
	}

	if len(indexes.UpProperties) != 1 {
		t.Errorf("Expected 1 UpProperty, got %d", len(indexes.UpProperties))
	}

	if indexes.UpProperties[3].Section.Header.Name != "Interface" {
		t.Errorf("Expected UpProperty section name 'Interface', got '%s'", indexes.UpProperties[0].Section.Header.Name)
	}

	if indexes.UpProperties[3].Property.Key.Name != "PreUp" {
		t.Errorf("Expected UpProperty key name 'PreUp', got '%s'", indexes.UpProperties[0].Property.Key.Name)
	}

	if indexes.UpProperties[3].Property.Value.Value != "iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT" {
		t.Errorf("Expected UpProperty value; %v", indexes.UpProperties[0].Property.Value.Value)
	}
}
