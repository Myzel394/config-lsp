package analyzer

import (
	testutils_test "config-lsp/handlers/fstab/test_utils"
	"testing"
)

func TestFSCKMultipleRoots(
	t *testing.T,
) {
	document := testutils_test.DocumentFromInput(t, `
UUID=12345678-1234-1234-1234-123456789012 /boot ext4 defaults 0 1
UUID=12345678-1234-1234-1234-123456789012 / btrfs defaults 0 1
UUID=12345678-1234-1234-1234-123456789012 /home ext4 defaults 0 2
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeFSCKField(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}

func TestFSCKBtrfsUsingRoot(
	t *testing.T,
) {
	document := testutils_test.DocumentFromInput(t, `
UUID=12345678-1234-1234-1234-123456789012 /boot btrfs defaults 0 1
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeFSCKField(ctx)

	if len(ctx.diagnostics) != 1 {
		t.Errorf("Expected 1 error, got %v", len(ctx.diagnostics))
	}
}
