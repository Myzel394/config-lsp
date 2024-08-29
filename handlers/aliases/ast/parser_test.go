package tree

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

	if !(len(parser.Aliases) == 1) {
		t.Fatalf("Expected 1 alias, got %v", len(parser.Aliases))
	}

	if !(parser.Aliases[0].Key.Value == "postmaster") {
		t.Fatalf("Expected key to be 'postmaster', got %v", parser.Aliases[1].Key.Value)
	}

	userValue := parser.Aliases[0].Values.Values[0].(AliasValueUser)
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

	if !(len(parser.Aliases) == 2) {
		t.Fatalf("Expected 2 aliases, got %v", len(parser.Aliases))
	}

	if !(parser.Aliases[0].Key.Value == "heinz") {
		t.Fatalf("Expected key to be 'heinz', got %v", parser.Aliases[1].Key.Value)
	}

	if !(parser.Aliases[1].Key.Value == "michel") {
		t.Fatalf("Expected key to be 'michel', got %v", parser.Aliases[1].Key.Value)
	}

	emailValue := parser.Aliases[1].Values.Values[0].(AliasValueEmail)

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

	if !(len(parser.Aliases) == 1) {
		t.Fatalf("Expected 1 alias, got %v", len(parser.Aliases))
	}

	if !(parser.Aliases[0].Key.Value == "luke") {
		t.Fatalf("Expected key to be 'luke', got %v", parser.Aliases[1].Key.Value)
	}

	includeValue := parser.Aliases[0].Values.Values[0].(AliasValueInclude)

	if !(includeValue.Path.Path == "/etc/other_aliases") {
		t.Fatalf("Expected path to be '/etc/other_aliases', got %v", includeValue.Path.Path)
	}

	if !(includeValue.Location.Start.Character == 6 && includeValue.Location.End.Character == 33) {
		t.Fatalf("Expected location to be 6-33, got %v-%v", includeValue.Location.Start.Character, includeValue.Location.End.Character)
	}

	if !(includeValue.Path.Location.Start.Character == 15 && includeValue.Path.Location.End.Character == 33) {
		t.Fatalf("Expected path location to be 15-33, got %v-%v", includeValue.Path.Location.Start.Character, includeValue.Path.Location.End.Character)
	}
}
