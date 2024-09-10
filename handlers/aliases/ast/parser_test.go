package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestSimpleExampleWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
postmaster: root
`)

	parser := NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(parser.Aliases.Size() == 1) {
		t.Fatalf("Expected 1 alias, got %v", parser.Aliases.Size())
	}

	rawEntry, _ := parser.Aliases.Get(uint32(0))
	entry := rawEntry.(*AliasEntry)
	if !(entry.Key.Value == "postmaster") {
		t.Fatalf("Expected key to be 'postmaster', got %v", entry.Key.Value)
	}

	userValue := entry.Values.Values[0].(AliasValueUser)
	if !(userValue.Value == "root") {
		t.Fatalf("Expected value to be 'root', got %v", userValue.Value)
	}

	if !(userValue.Location.Start.Line == 0) {
		t.Fatalf("Expected start line to be 0, got %v", userValue.Location.Start.Line)
	}
}

func TestMultipleValuesWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
heinz: root, goli
michel: raiks@example.com
`)
	parser := NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(parser.Aliases.Size() == 2) {
		t.Fatalf("Expected 2 aliases, got %v", parser.Aliases.Size())
	}

	rawEntry, _ := parser.Aliases.Get(uint32(0))
	entry := rawEntry.(*AliasEntry)
	if !(entry.Key.Value == "heinz") {
		t.Fatalf("Expected key to be 'heinz', got %v", entry.Key.Value)
	}

	rawEntry, _ = parser.Aliases.Get(uint32(1))
	entry = rawEntry.(*AliasEntry)
	if !(entry.Key.Value == "michel") {
		t.Fatalf("Expected key to be 'michel', got %v", entry.Key.Value)
	}

	emailValue := entry.Values.Values[0].(AliasValueEmail)

	if !(emailValue.Value == "raiks@example.com") {
		t.Fatalf("Expected value to be 'raiks@example.com', got %v", emailValue.Value)
	}
}

func TestIncludeWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
luke: :include:/etc/other_aliases
`)
	parser := NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	if !(parser.Aliases.Size() == 1) {
		t.Fatalf("Expected 1 alias, got %v", parser.Aliases.Size())
	}

	rawEntry, _ := parser.Aliases.Get(uint32(0))
	entry := rawEntry.(*AliasEntry)
	if !(entry.Key.Value == "luke") {
		t.Fatalf("Expected key to be 'luke', got %v", entry.Key.Value)
	}

	includeValue := entry.Values.Values[0].(AliasValueInclude)

	if !(includeValue.Path.Path == "/etc/other_aliases") {
		t.Fatalf("Expected path to be '/etc/other_aliases', got %v", includeValue.Path.Path)
	}

	if !(includeValue.Location.Start.Character == 6 && includeValue.Location.End.Character == 32) {
		t.Fatalf("Expected location to be 6-33, got %v-%v", includeValue.Location.Start.Character, includeValue.Location.End.Character)
	}

	if !(includeValue.Path.Location.Start.Character == 15 && includeValue.Path.Location.End.Character == 32) {
		t.Fatalf("Expected path location to be 15-33, got %v-%v", includeValue.Path.Location.Start.Character, includeValue.Path.Location.End.Character)
	}
}

func TestInvalidWithEmptyValuesWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
luke:
`)
	parser := NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) == 0 {
		t.Fatalf("Expected 1 error, got %v", errors)
	}

	if !(errors[0].Range.Start.Character == 5 && errors[0].Range.End.Character == 5) {
		t.Fatalf("Expected error to be at 6, got %v", errors[0].Range.Start.Character)
	}
}

func TestInvalidWithEmptyKeyWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
: root
`)
	parser := NewAliasesParser()
	errors := parser.Parse(input)

	if len(errors) == 0 {
		t.Fatalf("Expected 1 error, got %v", errors)
	}

	if !(errors[0].Range.Start.Character == 0 && errors[0].Range.End.Character == 0) {
		t.Fatalf("Expected error to be at 0, got %v", errors[0].Range.Start.Character)
	}
}
