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
		t.Fatalf("Expected no errors, got %v", ctx.diagnostics)
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
		t.Fatalf("Expected 2 errors (separator error and value error), got %v", ctx.diagnostics)
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

func TestUnknownOptionButIgnoredExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
IgnoreUnknown ThisOptionDoesNotExist
ThisOptionDoesNotExist okay
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(ctx.diagnostics))
	}

	if !(len(ctx.document.Indexes.UnknownOptions) == 0) {
		t.Errorf("Expected 0 unknown options, got %v", len(ctx.document.Indexes.UnknownOptions))
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
