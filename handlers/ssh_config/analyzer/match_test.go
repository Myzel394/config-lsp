package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestMatchInvalidAllArgument(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Match user lena all
`)
	ctx := &analyzerContext{
		document:    *d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeMatchBlocks(ctx)

	if !(len(ctx.diagnostics) == 1 && ctx.diagnostics[0].Range.Start.Line == 0) {
		t.Fatalf("Expected one error, got %v", ctx.diagnostics)
	}
}
