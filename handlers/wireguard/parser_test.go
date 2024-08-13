package wireguard

import (
	"strings"
	"testing"
)

func dedent(s string) string {
	return strings.TrimLeft(s, "\n")
}

func TestValidWildTestWorksFine(
	t *testing.T,
) {
	sample := dedent(`
[Interface]
PrivateKey = 1234567890
Address = 192.168.1.0/24

# I'm a comment
[Peer]
PublicKey = 1234567890
Endpoint = 1.2.3.4 ; I'm just a comment

[Peer]
PublicKey = 5555
	`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !(len(parser.CommentLines) == 1 && parser.CommentLines[0] == 4) {
		t.Fatalf("parseFromString failed to collect comment lines %v", parser.CommentLines)
	}

	if !((len(parser.Sections) == 3) && (*parser.Sections[0].Name == "Interface") && (*parser.Sections[1].Name == "Peer") && (*parser.Sections[2].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}

	if !(parser.Sections[0].StartLine == 0 && parser.Sections[0].EndLine == 2 && parser.Sections[1].StartLine == 5 && parser.Sections[1].EndLine == 7 && parser.Sections[2].StartLine == 9 && parser.Sections[2].EndLine == 10) {
		t.Fatalf("parseFromString: Invalid start and end lines %v", parser.Sections)
	}

	if !((len(parser.Sections[0].Properties) == 2) && (len(parser.Sections[1].Properties) == 2) && (len(parser.Sections[2].Properties) == 1)) {
		t.Fatalf("parseFromString: Invalid amount of properties %v", parser.Sections)
	}

	if !((parser.Sections[0].Properties[1].Key.Name == "PrivateKey") && (parser.Sections[0].Properties[2].Key.Name == "Address")) {
		t.Fatalf("parseFromString failed to collect properties of section 0 %v", parser.Sections[0].Properties)
	}

	if !((parser.Sections[1].Properties[6].Key.Name == "PublicKey") && (parser.Sections[1].Properties[7].Key.Name == "Endpoint")) {
		t.Fatalf("parseFromString failed to collect properties of section 1 %v", parser.Sections[1].Properties)
	}

	if !(parser.Sections[2].Properties[10].Key.Name == "PublicKey") {
		t.Fatalf("parseFromString failed to collect properties of section 2 %v", parser.Sections[2].Properties)
	}
}

func TestEmptySectionAtStartWorksFine(
	t *testing.T,
) {
	sample := dedent(`
[Interface]

[Peer]
PublicKey = 1234567890
`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Interface") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 0 && len(parser.Sections[1].Properties) == 1) {
		t.Fatalf("parseFromString failed to collect properties %v", parser.Sections)
	}
}

func TestEmptySectionAtEndWorksFine(
	t *testing.T,
) {
	sample := dedent(`
[Inteface]
PrivateKey = 1234567890

[Peer]
# Just sneaking in here, hehe
`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Inteface") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 1 && len(parser.Sections[1].Properties) == 0) {
		t.Fatalf("parseFromString failed to collect properties %v", parser.Sections)
	}

	if !(len(parser.CommentLines) == 1 && parser.CommentLines[0] == 4) {
		t.Fatalf("parseFromString failed to collect comment lines %v", parser.CommentLines)
	}
}

func TestEmptyFileWorksFine(
	t *testing.T,
) {
	sample := dedent(`
`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !(len(parser.Sections) == 0) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}
}

func TestPartialSectionWithNoPropertiesWorksFine(
	t *testing.T,
) {
	sample := dedent(`
[Inte

[Peer]
PublicKey = 1234567890
`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Inte") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 0 && len(parser.Sections[1].Properties) == 1) {
		t.Fatalf("parseFromString failed to collect properties: %v", parser.Sections)
	}

	if !(len(parser.CommentLines) == 0) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.CommentLines)
	}

	if !(parser.Sections[1].Properties[3].Key.Name == "PublicKey") {
		t.Fatalf("parseFromString failed to collect properties of section 1: %v", parser.Sections[1].Properties)
	}
}

func TestPartialSectionWithPropertiesWorksFine(
	t *testing.T,
) {
	sample := dedent(`
[Inte
PrivateKey = 1234567890

[Peer]
`)

	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Inte") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 1 && len(parser.Sections[1].Properties) == 0) {
		t.Fatalf("parseFromString failed to collect properties: %v", parser.Sections)
	}

	if !(parser.Sections[0].Properties[1].Key.Name == "PrivateKey") {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}
}

func TestFileWithOnlyComments(
	t *testing.T,
) {
	sample := dedent(`
# This is a comment
# Another comment
`)
	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(len(parser.Sections) == 0) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.CommentLines) == 2) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.CommentLines)
	}

	if !(parser.CommentLines[0] == 0 && parser.CommentLines[1] == 1) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.CommentLines)
	}
}

func TestMultipleSectionsNoProperties(
	t *testing.T,
) {
	sample := dedent(`
[Interface]
[Peer]
[Peer]
`)
	parser := wireguardParser{}
	errors := parser.parseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(len(parser.Sections) == 3) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	for _, section := range parser.Sections {
		if len(section.Properties) != 0 {
			t.Fatalf("parseFromString failed to collect properties: %v", section.Properties)
		}
	}
}
