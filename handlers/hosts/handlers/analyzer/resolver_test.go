package analyzer

import (
	"config-lsp/handlers/hosts/handlers/ast"
	"config-lsp/utils"
	"testing"
)

func TestResolverEntriesWorksWithNonOverlapping(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com
5.5.5.5 world.com
`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("PARER FAILED! Expected no errors, but got %v", errors)
	}

	resolver, errors := createResolverFromParser(parser)

	if len(errors) != 0 {
		t.Errorf("Expected no errors, but got %v", errors)
	}

	if len(resolver.Entries) != 2 {
		t.Errorf("Expected 2 entries, but got %v", len(resolver.Entries))
	}

	if !(resolver.Entries["hello.com"].IPv4Address.String() == "1.2.3.4") {
		t.Errorf("Expected hello.com to be 1.2.3.4, but got %v", resolver.Entries["hello.com"].IPv4Address)
	}

	if !(resolver.Entries["world.com"].IPv4Address.String() == "5.5.5.5") {
		t.Errorf("Expected world.com to be 5.5.5.5, but got %v", resolver.Entries["world.com"].IPv4Address)
	}

	if !(resolver.Entries["hello.com"].IPv6Address == nil) {
		t.Errorf("Expected hello.com to have no IPv6 address, but got %v", resolver.Entries["hello.com"].IPv6Address)
	}

	if !(resolver.Entries["world.com"].IPv6Address == nil) {
		t.Errorf("Expected world.com to have no IPv6 address, but got %v", resolver.Entries["world.com"].IPv6Address)
	}
}

func TestResolverEntriesWithSimpleOverlapping(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com
5.5.5.5 hello.com
`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("PARER FAILED! Expected no errors, but got %v", errors)
	}

	resolver, errors := createResolverFromParser(parser)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, but got %v", len(errors))
	}

	if len(resolver.Entries) != 1 {
		t.Errorf("Expected 1 entry, but got %v", len(resolver.Entries))
	}

	if !(resolver.Entries["hello.com"].IPv4Address.String() == "1.2.3.4") {
		t.Errorf("Expected hello.com to be 1.2.3.4, but got %v", resolver.Entries["hello.com"].IPv4Address)
	}
}

func TestResolverEntriesWithComplexOverlapping(
	t *testing.T,
) {
	input := utils.Dedent(`
1.2.3.4 hello.com test.com
5.5.5.5 check.com test.com
`)

	parser := ast.NewHostsParser()
	errors := parser.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("PARER FAILED! Expected no errors, but got %v", errors)
	}

	resolver, errors := createResolverFromParser(parser)

	if !(len(errors) == 1) {
		t.Errorf("Expected 1 error, but got %v", len(errors))
	}

	if len(resolver.Entries) != 3 {
		t.Errorf("Expected 3 entries, but got %v", len(resolver.Entries))
	}

	if !(resolver.Entries["hello.com"].IPv4Address.String() == "1.2.3.4") {
		t.Errorf("Expected hello.com to be 1.2.3.4, but got %v", resolver.Entries["hello.com"].IPv4Address)
	}

	if !(resolver.Entries["check.com"].IPv4Address.String() == "5.5.5.5") {
		t.Errorf("Expected check.com to be 5.5.5.5, but got %v", resolver.Entries["check.com"].IPv4Address)
	}

	if !(resolver.Entries["test.com"].IPv4Address.String() == "1.2.3.4") {
		t.Errorf("Expected test.com to have no IPv4 address, but got %v", resolver.Entries["test.com"].IPv4Address)
	}
}
