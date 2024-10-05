package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestIgnoreUnknownUnnecessary(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
IgnoreUnknown helloWorld
PermitRootLogin 'yes'
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}
	analyzeIgnoreUnknownHasNoUnnecessary(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}
