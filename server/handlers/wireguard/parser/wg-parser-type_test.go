package parser

import (
	"config-lsp/utils"
	"testing"
)

func TestGetLineTypeWorksCorrectly(
	t *testing.T,
) {
	sample := utils.Dedent(`
# A comment at the very top
Test=Hello

[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	parser := CreateWireguardParser()
	parser.ParseFromString(sample)

	lineType := parser.GetTypeByLine(0)
	if lineType != LineTypeComment {
		t.Fatalf("getTypeByLine: Expected line 0 to be a comment, but it is %v", lineType)
	}

	lineType = parser.GetTypeByLine(1)
	if lineType != LineTypeProperty {
		t.Fatalf("getTypeByLine: Expected line 1 to be a property, but it is %v", lineType)
	}

	lineType = parser.GetTypeByLine(2)
	if lineType != LineTypeEmpty {
		t.Fatalf("getTypeByLine: Expected line 2 to be empty, but it is %v", lineType)
	}

	lineType = parser.GetTypeByLine(3)
	if lineType != LineTypeHeader {
		t.Fatalf("getTypeByLine: Expected line 3 to be a header, but it is %v", lineType)
	}

	lineType = parser.GetTypeByLine(4)
	if lineType != LineTypeProperty {
		t.Fatalf("getTypeByLine: Expected line 4 to be a property, but it is %v", lineType)
	}

	lineType = parser.GetTypeByLine(12)
	if lineType != LineTypeComment {
		t.Fatalf("getTypeByLine: Expected line 12 to be a comment, but it is %v", lineType)
	}
}

func TestGetBelongingSectionWorksCorrectly(
	t *testing.T,
) {
	sample := utils.Dedent(`
# A comment at the very top
Test=Hello

[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	parser := CreateWireguardParser()
	parser.ParseFromString(sample)

	section := parser.GetBelongingSectionByLine(0)

	// Comment
	if section != nil {
		t.Fatalf("getBelongingSectionByLine: Expected line 0 to be in no section, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(1)

	if section != parser.Sections[1] {
		t.Fatalf("getBelongingSectionByLine: Expected line 1 to be in global section, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(2)
	if section != parser.Sections[1] {
		t.Fatalf("getBelongingSectionByLine: Expected line 2 to be in global section, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(3)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 3 to be in section Interface, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(4)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 4 to be in section Interface, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(6)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 6 to be in section Interface, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(10)
	if section != parser.Sections[3] {
		t.Fatalf("getBelongingSectionByLine: Expected line 10 to be in section Peer, but it is in %v", section)
	}

	section = parser.GetBelongingSectionByLine(12)

	// Comment
	if section != nil {
		t.Fatalf("getBelongingSectionByLine: Expected line 12 to be in no section, but it is in %v", section)
	}
}
