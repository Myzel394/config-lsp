package indexes

import (
	"config-lsp/handlers/bitcoin_conf/ast"
	"testing"
)

func TestSectionsInvalidExample(t *testing.T) {
	content := `
[main]
addnode = 10.0.0.1

[main]
addnode = 10.0.0.2
`

	config := ast.NewBTCConfig()
	errs := config.Parse(content)

	if len(errs) > 0 {
		t.Fatalf("Failed to parse config: %v", errs)
	}

	indexes, errs := CreateIndexes(config)

	if len(errs) == 0 {
		t.Fatalf("Expected errors for duplicate sections, got none")
	}

	if len(indexes.SectionsByName) != 1 {
		t.Fatalf("Expected 1 section in indexes, got %d", len(indexes.SectionsByName))
	}

	if indexes.SectionsByName["main"].Start.Line != 1 {
		t.Fatalf("Expected section 'main' to start at line 0, got %d", indexes.SectionsByName["main"].Start.Line)
	}
}
