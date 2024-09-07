package indexes

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/utils"
	"testing"
)

func TestComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
postmaster: alice, bob
alice: root
bob: root
`)
	parser := ast.NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Unexpected errors: %v", errors)
	}

	indexes, errors := CreateIndexes(parser)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got: %v", errors)
	}

	if !(len(indexes.Keys) == 3) {
		t.Errorf("Expected 3 keys, but got %v", len(indexes.Keys))
	}

	if !(len(indexes.UserOccurrences) == 3) {
		t.Errorf("Expected 3 user occurrences, but got %v", len(indexes.UserOccurrences))
	}

	if !(len(indexes.UserOccurrences["root"]) == 2) {
		t.Errorf("Expected 2 occurrences of root, but got %v", len(indexes.UserOccurrences["root"]))
	}

	if !(len(indexes.UserOccurrences["alice"]) == 1) {
		t.Errorf("Expected 1 occurrence of alice, but got %v", len(indexes.UserOccurrences["alice"]))
	}

	if !(len(indexes.UserOccurrences["bob"]) == 1) {
		t.Errorf("Expected 1 occurrence of bob, but got %v", len(indexes.UserOccurrences["bob"]))
	}
}
