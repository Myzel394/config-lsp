package testutils_test

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"
	"testing"
)

func DocumentFromInput(
	t *testing.T,
	content string,
) *sshdconfig.SSHDocument {
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

	return &sshdconfig.SSHDocument{
		Config:  c,
		Indexes: i,
	}
}
