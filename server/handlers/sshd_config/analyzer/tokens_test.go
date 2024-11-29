package analyzer

import (
	testutils_test "config-lsp/handlers/sshd_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestInvalidTokensForNonExisting(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ThisOptionDoesNotExist Hello%%World
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeTokens(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestInvalidTokensForExistingOption(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Tunnel Hello%%World
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeTokens(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestValidTokens(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
AuthorizedPrincipalsCommand Hello World %% and %d
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeTokens(ctx)

	if len(ctx.diagnostics) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(ctx.diagnostics))
	}
}
