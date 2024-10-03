package handlers

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"
	"testing"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSimpleFormattingExampleWorks(
	t *testing.T,
) {
	input := utils.Dedent(`
		PermitRootLogin    yes   
a b 
`)
	config := ast.NewSSHDConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Failed to parse SSHD config: %v", errors)
	}

	i, errors := indexes.CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Failed to create indexes: %v", errors)
	}

	d := sshdconfig.SSHDDocument{
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
				Line:      1,
				Character: protocol.UInteger(0),
			},
		},
		options,
	)

	if err != nil {
		t.Errorf("Failed to format document: %v", err)
	}

	_ = edits
}
