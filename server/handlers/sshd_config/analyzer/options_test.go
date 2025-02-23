package analyzer

import (
	testutils_test "config-lsp/handlers/sshd_config/test_utils"
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
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}

	if !(len(ctx.document.Indexes.UnknownOptions) == 1) {
		t.Errorf("Expected 1 unknown option, got %v", len(ctx.document.Indexes.UnknownOptions))
	}

	if !(ctx.document.Indexes.UnknownOptions[0].Option.Key.Value.Value == "ThisOptionDoesNotExist") {
		t.Errorf("Expected 'ThisOptionDoesNotExist', got %v", ctx.document.Indexes.UnknownOptions[0].Option.Key.Value.Value)
	}
}
