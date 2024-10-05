package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestDuplicateHostExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host example.com
	User root

Host example.com
	User test
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeHostBlock(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestDuplicateMultipleHostExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host example.com google.com 
	User root

Host test.com example.com
	User test
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeHostBlock(ctx)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestNonDuplicateHostExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host example.com
	User root

Host example2.com
	User test
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeHostBlock(ctx)

	if !(len(ctx.diagnostics) == 0) {
		t.Errorf("Expected 0 error, got %v", len(ctx.diagnostics))
	}
}
