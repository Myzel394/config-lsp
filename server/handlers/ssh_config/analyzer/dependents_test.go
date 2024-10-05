package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleDependentExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
CanonicalDomains test.com
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeDependents(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestSimpleDependentExistsExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
CanonicalizeHostname yes
CanonicalDomains test.com
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeDependents(ctx)

	if len(ctx.diagnostics) > 0 {
		t.Errorf("Expected no errors, got %v", len(ctx.diagnostics))
	}
}
