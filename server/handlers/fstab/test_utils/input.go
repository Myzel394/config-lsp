package testutils_test

import (
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/shared"
	"config-lsp/utils"
	"testing"
)

func DocumentFromInput(
	t *testing.T,
	content string,
) *shared.FstabDocument {
	input := utils.Dedent(content)
	c := ast.NewFstabConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	return &shared.FstabDocument{
		Config: c,
	}
}
