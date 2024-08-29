package analyzer

import (
	"config-lsp/handlers/hosts/handlers/ast"
	"config-lsp/utils"
	"net"
	"testing"
)

func TestValidSimpleExampleWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com
	`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, but got %v", errors)
	}

	if !(len(parser.Tree.Entries) == 1) {
		t.Errorf("Expected 1 entry, but got %v", len(parser.Tree.Entries))
	}

	if parser.Tree.Entries[0].IPAddress == nil {
		t.Errorf("Expected IP address to be present, but got nil")
	}

	if !(parser.Tree.Entries[0].IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be 1.2.3.4, but got %v", parser.Tree.Entries[0].IPAddress.Value)
	}

	if !(parser.Tree.Entries[0].Hostname.Value == "hello.com") {
		t.Errorf("Expected hostname to be hello.com, but got %v", parser.Tree.Entries[0].Hostname.Value)
	}

	if !(parser.Tree.Entries[0].Aliases == nil) {
		t.Errorf("Expected no aliases, but got %v", parser.Tree.Entries[0].Aliases)
	}

	if !(parser.Tree.Entries[0].Location.Start.Line == 0) {
		t.Errorf("Expected line to be 1, but got %v", parser.Tree.Entries[0].Location.Start.Line)
	}

	if !(parser.Tree.Entries[0].Location.Start.Character == 0) {
		t.Errorf("Expected start to be 0, but got %v", parser.Tree.Entries[0].Location.Start)
	}

	if !(parser.Tree.Entries[0].Location.End.Character == 17) {
		t.Errorf("Expected end to be 17, but got %v", parser.Tree.Entries[0].Location.End.Character)
	}

	if !(parser.Tree.Entries[0].IPAddress.Location.Start.Line == 0) {
		t.Errorf("Expected IP address line to be 1, but got %v", parser.Tree.Entries[0].IPAddress.Location.Start.Line)
	}

	if !(parser.Tree.Entries[0].IPAddress.Location.Start.Character == 0) {
		t.Errorf("Expected IP address start to be 0, but got %v", parser.Tree.Entries[0].IPAddress.Location.Start.Character)
	}

	if !(parser.Tree.Entries[0].IPAddress.Location.End.Character == 7) {
		t.Errorf("Expected IP address end to be 7, but got %v", parser.Tree.Entries[0].IPAddress.Location.End.Character)
	}

	if !(len(parser.CommentLines) == 0) {
		t.Errorf("Expected no comment lines, but got %v", len(parser.CommentLines))
	}
}

func TestValidComplexExampleWorks(
	t *testing.T,
) {
	input := utils.Dedent(`

# This is a comment
1.2.3.4 hello.com test.com # This is another comment
5.5.5.5 test.com
1.2.3.4 example.com check.com
`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, but got %v", errors)
	}

	if !(len(parser.Tree.Entries) == 3) {
		t.Errorf("Expected 3 entries, but got %v", len(parser.Tree.Entries))
	}

	if parser.Tree.Entries[2].IPAddress == nil {
		t.Errorf("Expected IP address to be present, but got nil")
	}

	if !(parser.Tree.Entries[2].IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be 1.2.3.4, but got %v", parser.Tree.Entries[2].IPAddress.Value)
	}

	if !(len(parser.CommentLines) == 1) {
		t.Errorf("Expected 1 comment line, but got %v", len(parser.CommentLines))
	}

	if !(utils.KeyExists(parser.CommentLines, 1)) {
		t.Errorf("Expected comment line 2 to exist, but it does not")
	}
}

func TestInvalidExampleWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4
	`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) == 0 {
		t.Errorf("Expected errors, but got none")
	}

	if !(len(parser.Tree.Entries) == 1) {
		t.Errorf("Expected 1 entries, but got %v", len(parser.Tree.Entries))
	}

	if !(len(parser.CommentLines) == 0) {
		t.Errorf("Expected no comment lines, but got %v", len(parser.CommentLines))
	}

	if !(parser.Tree.Entries[0].IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be nil, but got %v", parser.Tree.Entries[0].IPAddress)
	}

	if !(parser.Tree.Entries[0].Hostname == nil) {
		t.Errorf("Expected hostname to be nil, but got %v", parser.Tree.Entries[0].Hostname)
	}

	if !(parser.Tree.Entries[0].Aliases == nil) {
		t.Errorf("Expected aliases to be nil, but got %v", parser.Tree.Entries[0].Aliases)
	}
}
