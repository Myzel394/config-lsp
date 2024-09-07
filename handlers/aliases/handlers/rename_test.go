package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"
	"testing"
)

func TestRenameSimpleExample(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: alice
bob: root
support: alice, bob
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

	edits := RenameAlias(i, i.Keys["alice"], "amelie")

	if !(len(edits) == 3) {
		t.Errorf("Expected 2 edits, but got %v", len(edits))
	}

	if !(edits[0].Range.Start.Line == 0 && edits[0].Range.Start.Character == 0 && edits[0].Range.End.Line == 0 && edits[0].Range.End.Character == 5) {
		t.Errorf("Unexpected edit: %v", edits[0])
	}

	if !(edits[1].Range.Start.Line == 0 && edits[1].Range.Start.Character == 7 && edits[1].Range.End.Line == 0 && edits[1].Range.End.Character == 12) {
		t.Errorf("Unexpected edit: %v", edits[1])
	}

	if !(edits[2].Range.Start.Line == 2 && edits[2].Range.Start.Character == 9 && edits[2].Range.End.Line == 2 && edits[2].Range.End.Character == 14) {
		t.Errorf("Unexpected edit: %v", edits[2])
	}
}
