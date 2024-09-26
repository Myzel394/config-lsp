package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"
)

func TestSimpleInvalidQuotesExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
PermitRootLogin 'yes'
`)

	errors := analyzeQuotesAreValid(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestSingleQuotesKeyAndOptionExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
'Port' '22'
`)

	errors := analyzeQuotesAreValid(d)

	if !(len(errors) == 2) {
		t.Errorf("Expected 2 errors, got %v", len(errors))
	}
}

func TestSimpleUnclosedQuoteExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
PermitRootLogin "yes
`)

	errors := analyzeQuotesAreValid(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}

func TestIncompleteQuotesExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
"Port 
`)

	errors := analyzeQuotesAreValid(d)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, got %v", len(errors))
	}
}
