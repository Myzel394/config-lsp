package analyzer

import (
	"config-lsp/handlers/hosts/ast"
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
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(parser.Tree.Entries.Size() == 1) {
		t.Errorf("Expected 1 entry, but got %v", parser.Tree.Entries.Size())
	}

	rawEntry, found := parser.Tree.Entries.Get(uint32(0))
	if !found {
		t.Fatalf("Expected IP address to be present, but got nil")
	}

	entry := rawEntry.(*ast.HostsEntry)
	if !(entry.IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be 1.2.3.4, but got %v", entry.IPAddress.Value)
	}

	if !(entry.Hostname.Value == "hello.com") {
		t.Errorf("Expected hostname to be hello.com, but got %v", entry.Hostname.Value)
	}

	if !(entry.Aliases == nil) {
		t.Errorf("Expected no aliases, but got %v", entry.Aliases)
	}

	if !(entry.Location.Start.Line == 0) {
		t.Errorf("Expected line to be 1, but got %v", entry.Location.Start.Line)
	}

	if !(entry.Location.Start.Character == 0) {
		t.Errorf("Expected start to be 0, but got %v", entry.Location.Start)
	}

	if !(entry.Location.End.Character == 16) {
		t.Errorf("Expected end to be 17, but got %v", entry.Location.End.Character)
	}

	if !(entry.IPAddress.Location.Start.Line == 0) {
		t.Errorf("Expected IP address line to be 1, but got %v", entry.IPAddress.Location.Start.Line)
	}

	if !(entry.IPAddress.Location.Start.Character == 0) {
		t.Errorf("Expected IP address start to be 0, but got %v", entry.IPAddress.Location.Start.Character)
	}

	if !(entry.IPAddress.Location.End.Character == 6) {
		t.Errorf("Expected IP address end to be 7, but got %v", entry.IPAddress.Location.End.Character)
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
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(parser.Tree.Entries.Size() == 3) {
		t.Fatalf("Expected 3 entries, but got %v", parser.Tree.Entries.Size())
	}

	rawEntry, _ := parser.Tree.Entries.Get(uint32(2))
	entry := rawEntry.(*ast.HostsEntry)
	if entry.IPAddress == nil {
		t.Errorf("Expected IP address to be present, but got nil")
	}

	if !(entry.IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be 1.2.3.4, but got %v", entry.IPAddress.Value)
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
		t.Fatalf("Expected errors, but got none")
	}

	if !(parser.Tree.Entries.Size() == 1) {
		t.Errorf("Expected 1 entries, but got %v", parser.Tree.Entries.Size())
	}

	if !(len(parser.CommentLines) == 0) {
		t.Errorf("Expected no comment lines, but got %v", len(parser.CommentLines))
	}

	rawEntry, _ := parser.Tree.Entries.Get(uint32(0))
	entry := rawEntry.(*ast.HostsEntry)
	if !(entry.IPAddress.Value.String() == net.ParseIP("1.2.3.4").String()) {
		t.Errorf("Expected IP address to be nil, but got %v", entry.IPAddress)
	}

	if !(entry.Hostname == nil) {
		t.Errorf("Expected hostname to be nil, but got %v", entry.Hostname)
	}

	if !(entry.Aliases == nil) {
		t.Errorf("Expected aliases to be nil, but got %v", entry.Aliases)
	}
}
