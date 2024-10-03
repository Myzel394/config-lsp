package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
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
		t.Errorf("Expected 2 ctx.diagnostics, got %v", len(ctx.diagnostics))
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

func TestDependentOptionsExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Port 1234
CanonicalDomains example.com
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	option := d.FindOptionsByName("canonicaldomains")[0]
	checkIsDependent(ctx, option.Option.Key, option.Block)

	if !(len(ctx.diagnostics) == 1) {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestValidDependentOptionsExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Port 1234
CanonicalizeHostname yes
CanonicalDomains example.com
`)
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	option := d.FindOptionsByName("canonicaldomains")[0]
	checkIsDependent(ctx, option.Option.Key, option.Block)

	if len(ctx.diagnostics) > 0 {
		t.Errorf("Expected no errors, got %v", len(ctx.diagnostics))
	}
}
