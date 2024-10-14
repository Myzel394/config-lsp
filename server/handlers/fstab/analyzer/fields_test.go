package analyzer

import (
	testutils_test "config-lsp/handlers/fstab/test_utils"
	"testing"
)

func TestFieldsMissingMountPoint(t *testing.T) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeFieldAreFilled(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Fatalf("Expected 1 diagnostic, got %d", len(ctx.diagnostics))
	}
}

func TestValidExample(t *testing.T) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test /mnt/test ext4 defaults 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeFieldAreFilled(ctx)

	if len(ctx.diagnostics) != 0 {
		t.Fatalf("Expected 0 diagnostics, got %d", len(ctx.diagnostics))
	}
}
