package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"
)

func TestMatchInvalidAllArgument(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Match user lena all
`)

	errors := analyzeMatchBlocks(d)

	if !(len(errors) == 1 && errors[0].Range.Start.Line == 0) {
		t.Fatalf("Expected one error, got %v", errors)
	}
}
