package wireguard

import "testing"

func TestSimplePropertyInInterface(
	t *testing.T,
) {
	sample := dedent(`
[Interface]

`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions, err := parser.Sections[0].getCompletionsForEmptyLine()

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != len(interfaceOptions) {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", len(interfaceOptions), len(completions))
	}
}

func TestSimpleOneExistingPropertyInInterface(
	t *testing.T,
) {
	sample := dedent(`
[Interface]
PrivateKey = 1234567890

`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions, err := parser.Sections[0].getCompletionsForEmptyLine()

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	expected := len(interfaceOptions) - 1
	if len(completions) != expected {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", expected, len(completions))
	}
}

func TestEmptyRootCompletionsWork(
	t *testing.T,
) {
	sample := dedent(`
	`)

	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions := parser.getRootCompletionsForEmptyLine()

	if len(completions) != 2 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 2 completions, but got %v", len(completions))
	}
}

func TestInterfaceSectionRootCompletionsBeforeWork(
	t *testing.T,
) {
	sample := dedent(`

[Interface]
`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions := parser.getRootCompletionsForEmptyLine()

	if len(completions) != 1 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 1 completions, but got %v", len(completions))
	}
}

func TestInterfaceAndPeerSectionRootCompletionsWork(
	t *testing.T,
) {
	sample := dedent(`
[Interface]

[Peer]
`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions := parser.getRootCompletionsForEmptyLine()

	if len(completions) != 1 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 1 completions, but got %v", len(completions))
	}
}

func TestPropertyNoSepatorShouldCompleteSeparator(
	t *testing.T,
) {
	sample := dedent(`
[Interface]
DNS
`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions, err := parser.Sections[0].getCompletionsForPropertyLine(1, 3)

	if err == nil {
		t.Fatalf("getCompletionsForPropertyLine err is nil but should not be")
	}

	if len(completions) != 1 {
		t.Fatalf("getCompletionsForPropertyLine: Expected 1 completion, but got %v", len(completions))
	}

	if *completions[0].InsertText != " = " {
		t.Fatalf("getCompletionsForPropertyLine: Expected completion to be ' = ', but got '%v'", completions[0].Label)
	}
}

func TestPropertyNoSeparatorWithSpaceShouldCompleteSeparator(
	t *testing.T,
) {
	sample := dedent(`
[Interface]
DNS 
`)
	parser := createWireguardParser()
	parser.parseFromString(sample)

	completions, err := parser.Sections[0].getCompletionsForPropertyLine(1, 4)

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
