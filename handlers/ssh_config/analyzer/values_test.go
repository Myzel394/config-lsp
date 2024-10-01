package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestUnknownOptionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ThisOptionDoesNotExist okay
`)
	ctx := &analyzerContext{
		document:    *d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeValuesAreValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestUnknownOptionButIgnoredExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
IgnoreUnknown ThisOptionDoesNotExist
ThisOptionDoesNotExist okay
`)
	ctx := &analyzerContext{
		document:    *d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) > 0 {
		t.Errorf("Expected no errors, but got %v", len(ctx.diagnostics))
	}
}

func TestUnknownOptionIgnoredIsAfterDefinitionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ThisOptionDoesNotExist okay
IgnoreUnknown ThisOptionDoesNotExist
`)
	ctx := &analyzerContext{
		document:    *d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeValuesAreValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}
