package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestParserInvalidWithPort(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4:80 hello.com
`)
	parser := NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, but got %v", errors)
	}
}
