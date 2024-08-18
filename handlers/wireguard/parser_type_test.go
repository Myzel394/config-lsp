package wireguard

import "testing"

func TestGetLineTypeWorksCorrectly(
	t *testing.T,
) {
	sample := dedent(`
# A comment at the very top
Test=Hello

[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	parser := createWireguardParser()
	parser.parseFromString(sample)

	lineType := parser.getTypeByLine(0)
	if lineType != LineTypeComment {
		t.Fatalf("getTypeByLine: Expected line 0 to be a comment, but it is %v", lineType)
	}

	lineType = parser.getTypeByLine(1)
	if lineType != LineTypeProperty {
		t.Fatalf("getTypeByLine: Expected line 1 to be a property, but it is %v", lineType)
	}

	lineType = parser.getTypeByLine(2)
	if lineType != LineTypeEmpty {
		t.Fatalf("getTypeByLine: Expected line 2 to be empty, but it is %v", lineType)
	}

	lineType = parser.getTypeByLine(3)
	if lineType != LineTypeHeader {
		t.Fatalf("getTypeByLine: Expected line 3 to be a header, but it is %v", lineType)
	}

	lineType = parser.getTypeByLine(4)
	if lineType != LineTypeProperty {
		t.Fatalf("getTypeByLine: Expected line 4 to be a property, but it is %v", lineType)
	}

	lineType = parser.getTypeByLine(12)
	if lineType != LineTypeComment {
		t.Fatalf("getTypeByLine: Expected line 12 to be a comment, but it is %v", lineType)
	}
}

func TestGetBelongingSectionWorksCorrectly(
	t *testing.T,
) {
	sample := dedent(`
# A comment at the very top
Test=Hello

[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	parser := createWireguardParser()
	parser.parseFromString(sample)

	section := parser.getBelongingSectionByLine(0)

	// Comment
	if section != nil {
		t.Fatalf("getBelongingSectionByLine: Expected line 0 to be in no section, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(1)

	if section != parser.Sections[1] {
		t.Fatalf("getBelongingSectionByLine: Expected line 1 to be in global section, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(2)
	if section != parser.Sections[1] {
		t.Fatalf("getBelongingSectionByLine: Expected line 2 to be in global section, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(3)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 3 to be in section Interface, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(4)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 4 to be in section Interface, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(6)
	if section != parser.Sections[2] {
		t.Fatalf("getBelongingSectionByLine: Expected line 6 to be in section Interface, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(10)
	if section != parser.Sections[3] {
		t.Fatalf("getBelongingSectionByLine: Expected line 10 to be in section Peer, but it is in %v", section)
	}

	section = parser.getBelongingSectionByLine(12)

	// Comment
	if section != nil {
		t.Fatalf("getBelongingSectionByLine: Expected line 12 to be in no section, but it is in %v", section)
	}
}
