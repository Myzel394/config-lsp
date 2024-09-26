package analyzer

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"
	"testing"
)

func TestEmptyMatchBlocksMakesErrors(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Match User root
`)
	c := ast.NewSSHDConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	indexes, errors := indexes.CreateIndexes(*c)

	if len(errors) > 0 {
		t.Fatalf("Index error: %v", errors)
	}

	d := &sshdconfig.SSHDDocument{
		Config:  c,
		Indexes: indexes,
	}

	errors = analyzeMatchBlocks(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestContainsOnlyNegativeValues(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Match User !root,!admin
`)
	c := ast.NewSSHDConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	_, matchBlock := c.FindOption(uint32(1))
	errors = analyzeMatchValuesContainsPositiveValue(matchBlock.MatchValue.Entries[0].Values)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestEmptyMatchValues(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Match User
`)
	c := ast.NewSSHDConfig()

	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	i, errors := indexes.CreateIndexes(*c)

	if len(errors) > 0 {
		t.Fatalf("Index error: %v", errors)
	}

	d := &sshdconfig.SSHDDocument{
		Config:  c,
		Indexes: i,
	}

	errors = analyzeMatchBlocks(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestIncompleteMatchValues(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Match User 
`)
	c := ast.NewSSHDConfig()

	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	i, errors := indexes.CreateIndexes(*c)

	if len(errors) > 0 {
		t.Fatalf("Index error: %v", errors)
	}

	d := &sshdconfig.SSHDDocument{
		Config:  c,
		Indexes: i,
	}

	errors = analyzeMatchBlocks(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

