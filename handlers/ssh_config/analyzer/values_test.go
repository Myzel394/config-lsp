package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"
)

func TestUnknownOptionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ThisOptionDoesNotExist okay
`)

	errors := analyzeValuesAreValid(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestUnknownOptionButIgnoredExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
IgnoreUnknown ThisOptionDoesNotExist
ThisOptionDoesNotExist okay
`)

	errors := analyzeValuesAreValid(d)

	if len(errors) > 0 {
		t.Errorf("Expected no errors, but got %v", len(errors))
	}
}

func TestUnknownOptionIgnoredIsAfterDefinitionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ThisOptionDoesNotExist okay
IgnoreUnknown ThisOptionDoesNotExist
`)

	errors := analyzeValuesAreValid(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}
