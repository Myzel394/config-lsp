package analyzer

import (
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/indexes"
	"config-lsp/handlers/hosts/shared"
	"config-lsp/utils"
	"testing"
)

func TestWorksWithNonDoubleIPs(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com
5.5.5.5 world.com
1.2.3.5 foo.com
1.2.3.6 bar.com
`)

	parser := ast.NewHostsParser()
	i := indexes.NewHostsIndexes()
	document := hosts.HostsDocument{
		Parser:  &parser,
		Indexes: &i,
	}
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("PARER FAILED! Expected no errors, but got %v", errors)
	}

	errors = analyzeDoubleIPs(&document)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, but got %v", errors)
	}
}

func TestWorksWithDoubleIPs(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com
5.5.5.5 world.com
1.2.3.4 foo.com
`)

	parser := ast.NewHostsParser()
	i := indexes.NewHostsIndexes()
	document := hosts.HostsDocument{
		Parser:  &parser,
		Indexes: &i,
	}
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("PARER FAILED! Expected no errors, but got %v", errors)
	}

	errors = analyzeDoubleIPs(&document)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, but got %v", len(errors))
	}

	if !(errors[0].Range.Start.Line == 2) {
		t.Errorf("Expected error on line 2, but got %v", errors[0].Range.Start.Line)
	}

	if !(errors[0].Err.(shared.DuplicateIPDeclaration).AlreadyFoundAt == 0) {
		t.Errorf("Expected error on line 0, but got %v", errors[0].Err.(shared.DuplicateIPDeclaration).AlreadyFoundAt)
	}
}
