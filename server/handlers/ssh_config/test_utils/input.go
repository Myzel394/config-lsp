package testutils_test

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/indexes"
	"config-lsp/utils"
	"testing"
)

func DocumentFromInput(
	t *testing.T,
	content string,
) *sshconfig.SSHDocument {
	input := utils.Dedent(content)
	c := ast.NewSSHConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	i, errors := indexes.CreateIndexes(*c)

	if len(errors) > 0 {
		t.Fatalf("Index error: %v", errors)
	}

	return &sshconfig.SSHDocument{
		Config:  c,
		Indexes: i,
	}
}
