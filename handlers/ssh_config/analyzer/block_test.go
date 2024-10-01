package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"
)

func TestBlockEmptyBlock(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
Host *
`)

	errors := analyzeBlocks(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected an error, but got %v", len(errors))
	}
}
