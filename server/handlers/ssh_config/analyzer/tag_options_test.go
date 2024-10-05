package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestValidTagExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host test.com
	Tag auth

Match tagged auth
	User root
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeTagOptions(ctx)

	if len(ctx.diagnostics) > 0 {
		t.Errorf("Expected no errors, got %v", len(ctx.diagnostics))
	}
}

func TestTagBlockBeforeExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Match tagged auth
	User root

Host test.com
	Tag auth
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeTagOptions(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}
