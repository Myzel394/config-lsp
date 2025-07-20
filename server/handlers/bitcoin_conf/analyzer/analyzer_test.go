package analyzer

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/ast"
	"config-lsp/utils"
	"testing"
)

func TestAnalyzerValidExample(t *testing.T) {
	input := utils.Dedent(`
chain=main
addnode=10.0.0.1
	`)
	c := ast.NewBTCConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	document := &bitcoinconf.BTCDocument{
		Config: c,
	}

	errs := Analyze(document)

	if !(len(errs) == 0) {
		t.Fatalf("Expected 0 diagnostics, got %d: %v", len(errs), errs)
	}
}

func TestAnalyzerNonExistentPropertyExample(t *testing.T) {
	input := utils.Dedent(`
chain=main
nonexistent_property=value
	`)
	c := ast.NewBTCConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	document := &bitcoinconf.BTCDocument{
		Config: c,
	}

	errs := Analyze(document)

	if !(len(errs) > 0) {
		t.Fatalf("Expected diagnostics for nonexistent property, got none")
	}
}

func TestAnalyzerDuplicatePropertyExample(t *testing.T) {
	input := utils.Dedent(`
chain=main
chain=test
`)

	c := ast.NewBTCConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	document := &bitcoinconf.BTCDocument{
		Config: c,
	}

	errs := Analyze(document)

	if !(len(errs) > 0) {
		t.Fatalf("Expected diagnostics for property, got none")
	}
}

func TestAnalyzerChainProperty(t *testing.T) {
	input := utils.Dedent(`
chain=main
signet=1
`)

	c := ast.NewBTCConfig()
	errors := c.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Parse error: %v", errors)
	}

	document := &bitcoinconf.BTCDocument{
		Config: c,
	}

	errs := Analyze(document)

	if !(len(errs) > 0) {
		t.Fatalf("Expected diagnostics for property, got none")
	}
}
