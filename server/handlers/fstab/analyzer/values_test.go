package analyzer

import (
	testutils_test "config-lsp/handlers/fstab/test_utils"
	"testing"
)

func TestInvalidMountOptionsExample(
	t *testing.T,
) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test /mnt/test ext4 invalid 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) == 0 {
		t.Fatalf("Expected diagnostic, got %d", len(ctx.diagnostics))
	}
}

func TestExt4IsUsingBtrfsMountOption(
	t *testing.T,
) {
	document := testutils_test.DocumentFromInput(t, `
# Valid, but only for btrfs
LABEL=test /mnt/test ext4 subvolid=1 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) == 0 {
		t.Fatalf("Expected diagnostic, got %d", len(ctx.diagnostics))
	}
}

func TestValidBtrfsIsUsingBtrfsMountOption(
	t *testing.T,
) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test /mnt/test btrfs subvolid=1 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) != 0 {
		t.Fatalf("Expected diagnostic, got %d", len(ctx.diagnostics))
	}
}

func TestValidZFSArbitraryPropertyExample(t *testing.T) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test /mnt/test zfs my_arbitrary_user:property=1 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) != 0 {
		t.Fatalf("Expected no diagnostics, got %d", len(ctx.diagnostics))
	}
}

func TestInvalidBtrfsArbitraryPropertyExample(t *testing.T) {
	document := testutils_test.DocumentFromInput(t, `
LABEL=test /mnt/test btrfs my_arbitrary_user:property=1 0 0
`)

	ctx := &analyzerContext{
		document: document,
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) == 0 {
		t.Fatalf("Expected diagnostic, got %d", len(ctx.diagnostics))
	}
}
