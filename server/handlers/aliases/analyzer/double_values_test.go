package analyzer

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/utils"
	"testing"
)

func TestContainsDoubleUsers(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: root
master: alice, alice
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	// d := aliases.AliasesDocument{
	// 	Parser: &p,
	// }

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	errors = analyzeContainsNoDoubleValues(p)

	if !(len(errors) == 1) {
		t.Errorf("Expected errors, got none")
	}
}

func TestContainsDoubleEmails(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: root@localhost, some, noise, here, root@localhost
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	errors = analyzeContainsNoDoubleValues(p)

	if !(len(errors) == 1) {
		t.Errorf("Expected errors, got none")
	}
}

func TestContainsDoubleCommands(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: |echo, |test, |echo
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	errors = analyzeContainsNoDoubleValues(p)

	if !(len(errors) == 1) {
		t.Errorf("Expected errors, got none")
	}
}

func TestContainsDoubleErrors(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: error:450 Nonono, error:450 Some other message
root: error:450 Nonon, error:451 This is not okay
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	errors = analyzeContainsNoDoubleValues(p)

	if !(len(errors) == 1) {
		t.Errorf("Expected no errors, got %v", errors)
	}

	if !(errors[0].Range.Start.Line == 0 && errors[0].Range.End.Line == 0) {
		t.Errorf("Expected error to be on line 0, got %v-%v", errors[0].Range.Start.Line, errors[0].Range.End.Line)
	}
}

func TestComplexExampleContainsNoDoubleValues(
	t *testing.T,
) {
	input := utils.Dedent(`
alice: root@localhost, user@localhost
master: alice, root
noreply: error:450 Nonono
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	errors = analyzeContainsNoDoubleValues(p)

	if !(len(errors) == 0) {
		t.Errorf("Expected no errors, got %v", errors)
	}
}
