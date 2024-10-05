package analyzer

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/utils"
	"testing"
)

func TestWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
support: michael
marketing: john
support: jane
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	d := aliases.AliasesDocument{
		Parser: &p,
	}

	if len(errors) != 0 {
		t.Errorf("Expected no errors, got %v", errors)
	}

	errors = analyzeDoubleKeys(&d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", errors)
	}

	if d.Indexes == nil {
		t.Errorf("Expected indexes to be set")
	}
}

func TestValidWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
support: michael
marketing: john
supportgroup: jane
suppor: jane
`)
	p := ast.NewAliasesParser()
	errors := p.Parse(input)

	d := aliases.AliasesDocument{
		Parser: &p,
	}

	if len(errors) != 0 {
		t.Errorf("Expected no errors, got %v", errors)
	}

	errors = analyzeDoubleKeys(&d)

	if !(len(errors) == 0) {
		t.Errorf("Expected 0 errors, got %v", errors)
	}

	if d.Indexes == nil {
		t.Errorf("Expected indexes to be set")
	}

	if d.Indexes.Keys["support"] == nil {
		t.Errorf("Expected support to be in indexes")
	}
}
