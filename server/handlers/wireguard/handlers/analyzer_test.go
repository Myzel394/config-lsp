package handlers

import (
	"config-lsp/handlers/wireguard/parser"
	"config-lsp/utils"
	"testing"
)

func TestMultipleIntefaces(t *testing.T) {
	content := utils.Dedent(`
[Interface]
PrivateKey = abc

[Interface]
PrivateKey = def
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(content)

	diagnostics := Analyze(p)

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}

func TestInvalidValue(t *testing.T) {
	content := utils.Dedent(`
[Interface]
DNS = nope
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(content)

	diagnostics := Analyze(p)

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}

func TestDuplicateProperties(t *testing.T) {
	content := utils.Dedent(`
[Interface]
PrivateKey = abc
DNS = 1.1.1.1
PrivateKey = def
`)
	p := parser.CreateWireguardParser()
	p.ParseFromString(content)

	diagnostics := Analyze(p)

	if len(diagnostics) == 0 {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}
