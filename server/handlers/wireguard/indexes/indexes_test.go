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
	errs := config.Parse(content)

	if len(errs) > 0 {
		t.Fatalf("Failed to parse config: %v", errs)
	}

	indexes, errs := CreateIndexes(config)

	if len(errs) > 0 {
		t.Fatalf("Unexpected errors while creating indexes: %v", errs)
	}

	if !(indexes.AsymmetricRules[config.Sections[0]].PreMissing == true && indexes.AsymmetricRules[config.Sections[0]].PostMissing == false) {
		t.Errorf("Expected asymmetric rules for section '%s' to be PreMissing: true, PostMissing: false, got PreMissing: %v, PostMissing: %v", config.Sections[0].Header.Name, indexes.AsymmetricRules[config.Sections[0]].PreMissing, indexes.AsymmetricRules[config.Sections[0]].PostMissing)
	}
}
