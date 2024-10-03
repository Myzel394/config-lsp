package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ProxyCommand hello

User root
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) != 0 {
		t.Fatalf("Expected no errors, got %v", ctx.diagnostics)
	}
}

func TestOptionEmpty(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ProxyCommand 

User root
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Fatalf("Expected 1 error, got %v", ctx.diagnostics)
	}
}

func TestNoSeparator(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
"ProxyCommand""hello"

User root
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Fatalf("Expected 1 error, got %v", ctx.diagnostics)
	}
}

func TestEmptyMatch(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
User root

Host example.com
	User test

Match
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) != 2 {
		t.Fatalf("Expected 1 error, got %v", ctx.diagnostics)
	}
}

func TestEmptyWithSeparatorMatch(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
User root

Host example.com
	User test

Match 
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Fatalf("Expected 1 error, got %v", ctx.diagnostics)
	}
}
