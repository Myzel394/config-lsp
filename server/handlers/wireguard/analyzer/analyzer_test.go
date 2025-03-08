package analyzer

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
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
	d := &wireguard.WGDocument{
		Config: ast.NewWGConfig(),
	}
	d.Config.Parse(content)

	diagnostics := Analyze(d)

	if !(len(diagnostics) > 0) {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}

func TestInvalidValue(t *testing.T) {
	content := utils.Dedent(`
[Interface]
DNS = nope
`)
	d := &wireguard.WGDocument{
		Config: ast.NewWGConfig(),
	}
	d.Config.Parse(content)

	diagnostics := Analyze(d)

	if !(len(diagnostics) > 0) {
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

	d := &wireguard.WGDocument{
		Config: ast.NewWGConfig(),
	}
	d.Config.Parse(content)

	diagnostics := Analyze(d)

	if !(len(diagnostics) > 0) {
		t.Errorf("Expected diagnostic errors, got %d", len(diagnostics))
	}
}
