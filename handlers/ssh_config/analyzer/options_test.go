package analyzer

import (
	testutils_test "config-lsp/handlers/ssh_config/test_utils"
	"testing"
)

func TestSimpleExample(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ProxyCommand hello

User root
`)

	errors := analyzeStructureIsValid(d)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}
}

func TestOptionEmpty(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
ProxyCommand 

User root
`)
	errors := analyzeStructureIsValid(d)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %v", errors)
	}
}

func TestNoSeparator(
	t *testing.T,
) {
	d := testutils_test.DocumentFromInput(t, `
"ProxyCommand""hello"

User root
`)
	errors := analyzeStructureIsValid(d)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %v", errors)
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
	errors := analyzeStructureIsValid(d)

	if len(errors) != 2 {
		t.Fatalf("Expected 1 error, got %v", errors)
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
	errors := analyzeStructureIsValid(d)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %v", errors)
	}
}
