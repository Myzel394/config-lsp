package commands

import "testing"

func TestAreAliasesCommandsAvailable(
	t *testing.T,
) {
	if !IsNewAliasesCommandAvailable() {
		t.Skip("Aliases tools not available")
	}
}
