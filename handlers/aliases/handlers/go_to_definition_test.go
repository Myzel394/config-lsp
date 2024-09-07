package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"
	"testing"
)

func TestGoToDefinitionSimpleExample(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: root
bob: root
steve: alice@example.com, bob
david: alice
`)
	parser := ast.NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Unexpected errors: %v", errors)
	}

	i, errors := indexes.CreateIndexes(parser)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got: %v", errors)
	}

	rawEntry, _ := parser.Aliases.Get(uint32(3))
	entry := rawEntry.(*ast.AliasEntry)
	rawValue := entry.Values.Values[0]
	value := rawValue.(ast.AliasValueUser)

	locations := GetDefinitionLocationForValue(
		i,
		value,
		"file:///etc/aliases",
	)

	if !(len(locations) == 1) {
		t.Errorf("Expected 1 location, but got %v", len(locations))
	}

	if !(locations[0].URI == "file:///etc/aliases") {
		t.Errorf("Unexpected location: %v", locations[0])
	}

	if !(locations[0].Range.Start.Line == 0 && locations[0].Range.Start.Character == 0 && locations[0].Range.End.Line == 0 && locations[0].Range.End.Character == 5) {
		t.Errorf("Unexpected location: %v", locations[0])
	}
}
