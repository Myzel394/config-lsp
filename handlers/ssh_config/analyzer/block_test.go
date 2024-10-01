package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestBlockEmptyBlock(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host *
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeBlocks(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected an error, but got %v", len(ctx.diagnostics))
	}
}
