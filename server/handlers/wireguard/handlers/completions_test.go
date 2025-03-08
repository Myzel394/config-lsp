package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimplePropertyInInterface(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]

`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      1,
				Character: 0,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Errorf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != len(fields.InterfaceOptions) {
		t.Errorf("getCompletionsForEmptyLine: Expected %v completions, but got %v", len(fields.InterfaceOptions), len(completions))
	}
}

func TestSimpleOneExistingPropertyInInterface(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
PrivateKey = 1234567890

`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      2,
				Character: 0,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	expected := len(fields.InterfaceOptions) - 1
	if len(completions) != expected {
		t.Fatalf("getCompletionsForEmptyLine: Expected %v completions, but got %v", expected, len(completions))
	}
}

func TestEmptyCompletions(
	t *testing.T,
) {
	sample := utils.Dedent(`

`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      0,
				Character: 0,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Fatalf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != 2 {
		t.Fatalf("getRootCompletionsForEmptyLine: Expected 2 completions, but got %v", len(completions))
	}
}

func TestIncompletePropertyCompletions(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Peer]
Add
`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      1,
				Character: 3,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Errorf("getCompletionsForEmptyLine failed with error: %v", err)
	}

	if len(completions) != len(fields.PeerOptions) {
		t.Errorf("getRootCompletionsForEmptyLine: Expected 1 completions, but got %v", len(completions))
	}
}

func TestPropertyBeforeLineIsEmpty(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
DNS = 1.1.1.1


`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      3,
				Character: 0,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Errorf("getCompletionsForPropertyLine failed with error: %v", err)
	}

	if len(completions) != len(fields.InterfaceOptions)+1-1 {
		t.Errorf("getCompletionsForPropertyLine: Expected completions, but got %v", len(completions))
	}
}

func TestPropertyValueCompletions(
	t *testing.T,
) {
	sample := utils.Dedent(`
[Interface]
Table = 
`)
	p := ast.NewWGConfig()
	parseErrors := p.Parse(sample)

	if len(parseErrors) > 0 {
		t.Fatalf("Parser failed with error %v", parseErrors)
	}

	d := &wireguard.WGDocument{
		Config: p,
	}

	params := &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			Position: protocol.Position{
				Line:      1,
				Character: 8,
			},
		},
	}
	completions, err := SuggestCompletions(d, params)

	if err != nil {
		t.Errorf("getCompletionsForPropertyLine failed with error: %v", err)
	}

	if !(len(completions) == 2) {
		t.Errorf("getCompletionsForPropertyLine: Expected completions, but got %v", len(completions))
	}
}
