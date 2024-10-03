package analyzer

import (
	testutils_test "config-lsp/handlers/sshd_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleInvalidQuotesExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
PermitRootLogin 'yes'
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeQuotesAreValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestSingleQuotesKeyAndOptionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
'Port' '22'
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeQuotesAreValid(ctx)

	if !(len(ctx.diagnostics) == 2) {
		t.Errorf("Expected 2 errors, got %v", len(ctx.diagnostics))
	}
}

func TestSimpleUnclosedQuoteExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
PermitRootLogin "yes
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeQuotesAreValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestIncompleteQuotesExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
"Port 
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeQuotesAreValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}
