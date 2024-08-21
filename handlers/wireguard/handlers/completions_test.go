package handlers

import (
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/parser"
	"config-lsp/utils"
	"testing"
)

func TestSimplePropertyInInterface(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]

`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, err := GetCompletionsForSectionEmptyLine(*p.Sections[0])

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != len(fields.InterfaceOptions) {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", len(fields.InterfaceOptions), len(completions))
	}
}

func TestSimpleOneExistingPropertyInInterface(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
PrivateKey = 1234567890

`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, err := GetCompletionsForSectionEmptyLine(*p.Sections[0])

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	expected := len(fields.InterfaceOptions) - 1
	if len(completions) != expected {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", expected, len(completions))
	}
}

func TestEmptyRootCompletionsWork(
	t *testing.T,
) {
	sample := utils.Dedent(`
	`)

	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, _ := GetRootCompletionsForEmptyLine(p)

	if len(completions) != 2 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 2 completions, but got %v", len(completions))
	}
}

func TestInterfaceSectionRootCompletionsBeforeWork(
	t *testing.T,
) {
	sample := utils.Dedent(`

[Interface]
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, _ := GetRootCompletionsForEmptyLine(p)

	if len(completions) != 1 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 1 completions, but got %v", len(completions))
	}
}

func TestInterfaceAndPeerSectionRootCompletionsWork(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]

[Peer]
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, _ := GetRootCompletionsForEmptyLine(p)

	if len(completions) != 1 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 1 completions, but got %v", len(completions))
	}
}

func TestPropertyNoSepatorShouldCompleteSeparator(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, err := GetCompletionsForSectionPropertyLine(*p.Sections[0], 1, 3)

	if err == nil {
		t.Fatalf("getCompletionsForPropertyLine err is nil but should not be")
	}

	if len(completions) != 1 {
		t.Fatalf("getCompletionsForPropertyLine: Expected 1 completion, but got %v", len(completions))
	}

	if *completions[0].InsertText != "DNS = " {
		t.Fatalf("getCompletionsForPropertyLine: Expected completion to be 'DNS = ', but got '%v'", completions[0].Label)
	}
}

func TestPropertyNoSeparatorWithSpaceShouldCompleteSeparator(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS 
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, err := GetCompletionsForSectionPropertyLine(*p.Sections[0], 1, 4)

	if err == nil {
		t.Fatalf("getCompletionsForPropertyLine err is nil but should not be")
	}

	if len(completions) != 1 {
		t.Fatalf("getCompletionsForPropertyLine: Expected 1 completion, but got %v", len(completions))
	}

	if *completions[0].InsertText != "= " {
		t.Fatalf("getCompletionsForPropertyLine: Expected completion to be '= ', but got '%v'", completions[0].Label)
	}
}

func TestHeaderButNoProperty(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]

`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(sample)

	completions, err := GetCompletionsForSectionEmptyLine(*p.Sections[0])

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != len(fields.InterfaceOptions) {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", len(fields.InterfaceOptions), len(completions))
	}
}
