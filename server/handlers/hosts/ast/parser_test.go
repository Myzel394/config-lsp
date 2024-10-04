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

func TestParserValidComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com alias.com example.com
1.2.3.5 hello1.com alias1.com example1.com
192.168.1.1 goodbye.com
`)
	parser := NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(parser.Tree.Entries.Size() == 3) {
		t.Fatalf("Expected 3 entries, but got %v", parser.Tree.Entries.Size())
	}

	rawEntry, _ := parser.Tree.Entries.Get(uint32(0))
	entry := rawEntry.(*HostsEntry)

	if !(entry.IPAddress.Value.String() == "1.2.3.4") {
		t.Errorf("Expected IP address to be 1.2.3.4, but got %v", entry.IPAddress.Value)
	}

	if !(entry.Hostname.Value == "hello.com") {
		t.Errorf("Expected hostname to be hello.com, but got %v", entry.Hostname.Value)
	}
}
