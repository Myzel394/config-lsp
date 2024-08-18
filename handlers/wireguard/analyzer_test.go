package wireguard

import "testing"

func TestMultipleIntefaces(t *testing.T) {
	content := dedent(`
[Interface]
PrivateKey = abc

[Interface]
PrivateKey = def
`)
	parser := createWireguardParser()
	parser.parseFromString(content)

	diagnostics := parser.analyze()

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}

func TestInvalidValue(t *testing.T) {
	content := dedent(`
[Interface]
DNS = nope
`)
	parser := createWireguardParser()
	parser.parseFromString(content)

	diagnostics := parser.analyze()

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}

func TestDuplicateProperties(t *testing.T) {
	content := dedent(`
[Interface]
PrivateKey = abc
DNS = 1.1.1.1
PrivateKey = def
`)
	parser := createWireguardParser()
	parser.parseFromString(content)

	diagnostics := parser.analyze()

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}
