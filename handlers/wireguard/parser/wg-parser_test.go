package parser

import (
	"config-lsp/utils"
	"testing"

	"github.com/k0kubun/pp"
)

func TestValidWildTestWorksFine(
	t *testing.T,
) {
	sample := utils.Dedent(`
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

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !(len(parser.commentLines) == 1 && utils.KeyExists(parser.commentLines, 4)) {
		t.Fatalf("parseFromString failed to collect comment lines %v", parser.commentLines)
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

	// Check if line indexes are correct
	if !(parser.linesIndexes[0].Type == LineTypeHeader &&
		parser.linesIndexes[1].Type == LineTypeProperty &&
		parser.linesIndexes[2].Type == LineTypeProperty &&
		parser.linesIndexes[3].Type == LineTypeEmpty &&
		parser.linesIndexes[4].Type == LineTypeComment &&
		parser.linesIndexes[5].Type == LineTypeHeader &&
		parser.linesIndexes[6].Type == LineTypeProperty &&
		parser.linesIndexes[7].Type == LineTypeProperty &&
		parser.linesIndexes[8].Type == LineTypeEmpty &&
		parser.linesIndexes[9].Type == LineTypeHeader &&
		parser.linesIndexes[10].Type == LineTypeProperty) {
		pp.Println(parser.linesIndexes)
		t.Fatal("parseFromString: Invalid line indexes")
	}
}

func TestEmptySectionAtStartWorksFine(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]

[Peer]
PublicKey = 1234567890
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

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
	sample := utils.Dedent(`
[Inteface]
PrivateKey = 1234567890

[Peer]
# Just sneaking in here, hehe
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Inteface") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 1 && len(parser.Sections[1].Properties) == 0) {
		t.Fatalf("parseFromString failed to collect properties %v", parser.Sections)
	}

	if !(len(parser.commentLines) == 1 && utils.KeyExists(parser.commentLines, 4)) {
		t.Fatalf("parseFromString failed to collect comment lines %v", parser.commentLines)
	}
}

func TestEmptyFileWorksFine(
	t *testing.T,
) {
	sample := utils.Dedent(`
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !(len(parser.Sections) == 1) {
		t.Fatalf("parseFromString failed to collect sections %v", parser.Sections)
	}
}

func TestPartialSectionWithNoPropertiesWorksFine(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Inte

[Peer]
PublicKey = 1234567890
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error %v", errors)
	}

	if !((len(parser.Sections) == 2) && (*parser.Sections[0].Name == "Inte") && (*parser.Sections[1].Name == "Peer")) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 0 && len(parser.Sections[1].Properties) == 1) {
		t.Fatalf("parseFromString failed to collect properties: %v", parser.Sections)
	}

	if !(len(parser.commentLines) == 0) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.commentLines)
	}

	if !(parser.Sections[1].Properties[3].Key.Name == "PublicKey") {
		t.Fatalf("parseFromString failed to collect properties of section 1: %v", parser.Sections[1].Properties)
	}
}

func TestPartialSectionWithPropertiesWorksFine(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Inte
PrivateKey = 1234567890

[Peer]
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

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
	sample := utils.Dedent(`
# This is a comment
# Another comment
`)
	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(len(parser.Sections) == 1) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.commentLines) == 2) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.commentLines)
	}

	if !(utils.KeyExists(parser.commentLines, 0) && utils.KeyExists(parser.commentLines, 1)) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.commentLines)
	}
}

func TestMultipleSectionsNoProperties(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
[Peer]
[Peer]
`)

	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

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

func TestWildTest1WorksCorrectly(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS=1.1.1.1


`)
	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(len(parser.Sections) == 1) {
		t.Fatalf("parseFromString failed to collect sections: %v", parser.Sections)
	}

	if !(len(parser.Sections[0].Properties) == 1) {
		t.Fatalf("parseFromString failed to collect properties: %v", parser.Sections[0].Properties)
	}

	if !(parser.Sections[0].Properties[1].Key.Name == "DNS") {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}

	if !(parser.Sections[0].Properties[1].Value.Value == "1.1.1.1") {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}

	if !(len(parser.commentLines) == 0) {
		t.Fatalf("parseFromString failed to collect comment lines: %v", parser.commentLines)
	}

	if !(parser.Sections[0].StartLine == 0 && parser.Sections[0].EndLine == 1) {
		t.Fatalf("parseFromString: Invalid start and end lines %v", parser.Sections)
	}
}

func TestPartialKeyWorksCorrectly(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS
`)
	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(parser.Sections[0].Properties[1].Key.Name == "DNS") {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}

	if !(parser.Sections[0].Properties[1].Separator == nil) {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}
}

func TestPartialValueWithSeparatorWorksCorrectly(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS=
`)
	parser := CreateWireguardParser()
	errors := parser.ParseFromString(sample)

	if len(errors) > 0 {
		t.Fatalf("parseFromString failed with error: %v", errors)
	}

	if !(parser.Sections[0].Properties[1].Value == nil) {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}

	if !(parser.Sections[0].Properties[1].Separator != nil) {
		t.Fatalf("parseFromString failed to collect properties of section 0: %v", parser.Sections[0].Properties)
	}
}
