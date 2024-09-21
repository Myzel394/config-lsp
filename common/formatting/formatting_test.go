package formatting

import (
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleTabExampleWithTabOptions(
	t *testing.T,
) {
	template := FormatTemplate("%s/t%s")

	options := protocol.FormattingOptions{
		"tabSize":      float64(4),
		"insertSpaces": false,
	}

	result := template.Format(options, "PermitRootLogin", "yes")
	expected := "PermitRootLogin\tyes"

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestSimpleTabExampleWithSpaceOptions(
	t *testing.T,
) {
	template := FormatTemplate("%s/t%s")

	options := protocol.FormattingOptions{
		"tabSize":      float64(4),
		"insertSpaces": true,
	}

	result := template.Format(options, "PermitRootLogin", "yes")
	expected := "PermitRootLogin    yes"

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestSimpleExampleWhiteSpaceAtEndShouldTrim(
	t *testing.T,
) {
	template := FormatTemplate("%s/t%s")

	options := protocol.FormattingOptions{
		"tabSize":                float64(4),
		"insertSpaces":           false,
		"trimTrailingWhitespace": true,
	}

	result := template.Format(options, "PermitRootLogin", "yes    ")
	expected := "PermitRootLogin\tyes"

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestSimpleExampleWhiteSpaceAtEndShouldNOTTrim(
	t *testing.T,
) {
	template := FormatTemplate("%s/t%s")

	options := protocol.FormattingOptions{
		"tabSize":                float64(4),
		"insertSpaces":           false,
		"trimTrailingWhitespace": false,
	}

	result := template.Format(options, "PermitRootLogin", "yes    ")
	expected := "PermitRootLogin\tyes    "

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
