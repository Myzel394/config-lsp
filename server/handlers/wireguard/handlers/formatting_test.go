package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/indexes"
	"config-lsp/utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleFormattingExample(t *testing.T) {
	input := utils.Dedent(`
[Interface]
DNS = 9.9.9.9, 1.1.1.1
`)
	config := ast.NewWGConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Failed to parse WireGuard config: %v", errors)
	}

	i, errors := indexes.CreateIndexes(config)

	if len(errors) > 0 {
		t.Fatalf("Failed to create indexes: %v", errors)
	}

	d := wireguard.WGDocument{
		Config:  config,
		Indexes: i,
	}

	options := protocol.FormattingOptions{}
	options["insertSpaces"] = true
	options["tabSize"] = float64(4)

	edits, err := FormatDocument(
		&d,
		protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: protocol.UInteger(0),
			},
			End: protocol.Position{
				Line:      2,
				Character: protocol.UInteger(0),
			},
		},
		options,
	)

	if err != nil {
		t.Errorf("Failed to format document: %v", err)
	}

	if !(len(edits) == 1 && edits[0].NewText == `DNS = "9.9.9.9, 1.1.1.1"` && edits[0].Range.Start.Line == 1 && edits[0].Range.Start.Character == 0 && edits[0].Range.End.Line == 1 && edits[0].Range.End.Character == 22) {
		t.Errorf("Unexpected edits: %v", edits)
	}
}

func TestFormattingAlreadySurroundedQuotesExample(t *testing.T) {
	input := utils.Dedent(`
[Interface]
DNS = "9.9.9.9, 1.1.1.1"
`)
	config := ast.NewWGConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Failed to parse WireGuard config: %v", errors)
	}

	i, errors := indexes.CreateIndexes(config)

	if len(errors) > 0 {
		t.Fatalf("Failed to create indexes: %v", errors)
	}

	d := wireguard.WGDocument{
		Config:  config,
		Indexes: i,
	}

	options := protocol.FormattingOptions{}
	options["insertSpaces"] = true
	options["tabSize"] = float64(4)

	edits, err := FormatDocument(
		&d,
		protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: protocol.UInteger(0),
			},
			End: protocol.Position{
				Line:      2,
				Character: protocol.UInteger(0),
			},
		},
		options,
	)

	if err != nil {
		t.Errorf("Failed to format document: %v", err)
	}

	if !(len(edits) == 1 && edits[0].NewText == `DNS = "9.9.9.9, 1.1.1.1"` && edits[0].Range.Start.Line == 1 && edits[0].Range.Start.Character == 0 && edits[0].Range.End.Line == 1 && edits[0].Range.End.Character == 23) {
		t.Errorf("Unexpected edits: %v", edits)
	}
}

func TestWithinQuotesExample(t *testing.T) {
	input := utils.Dedent(`
[Interface]
DNS = 9.9.9.9, "1234:5678::1", 1.1.1.1
`)
	config := ast.NewWGConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Failed to parse WireGuard config: %v", errors)
	}

	i, errors := indexes.CreateIndexes(config)

	if len(errors) > 0 {
		t.Fatalf("Failed to create indexes: %v", errors)
	}

	d := wireguard.WGDocument{
		Config:  config,
		Indexes: i,
	}

	options := protocol.FormattingOptions{}
	options["insertSpaces"] = true
	options["tabSize"] = float64(4)

	edits, err := FormatDocument(
		&d,
		protocol.Range{
			Start: protocol.Position{
				Line:      0,
				Character: protocol.UInteger(0),
			},
			End: protocol.Position{
				Line:      2,
				Character: protocol.UInteger(0),
			},
		},
		options,
	)

	if err != nil {
		t.Errorf("Failed to format document: %v", err)
	}

	if !(len(edits) == 1 && edits[0].NewText == `DNS = "9.9.9.9, \"1234:5678::1\", 1.1.1.1"` && edits[0].Range.Start.Line == 1 && edits[0].Range.Start.Character == 0 && edits[0].Range.End.Line == 1 && edits[0].Range.End.Character == 38) {
		t.Errorf("Unexpected edits: %v", edits)
	}
}
