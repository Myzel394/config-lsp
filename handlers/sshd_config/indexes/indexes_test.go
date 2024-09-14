package indexes

import (
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/utils"
	"testing"
)

func TestSimpleExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
RootLogin yes
PermitRootLogin no
`)
	config := ast.NewSSHConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Unexpected errors: %v", errors)
	}

	indexes, errors := CreateIndexes(*config)

	if !(len(errors) == 1) {
		t.Fatalf("Expected 1 error, but got %v", len(errors))
	}

	indexEntry := SSHIndexKey{
		Option:     "PermitRootLogin",
		MatchBlock: nil,
	}
	if !(indexes.OptionsPerRelativeKey[indexEntry][0].Value == "PermitRootLogin yes" && indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line == 0) {
		t.Errorf("Expected 'PermitRootLogin yes', but got %v", indexes.OptionsPerRelativeKey[indexEntry][0].Value)
	}

	indexEntry = SSHIndexKey{
		Option:     "RootLogin",
		MatchBlock: nil,
	}
	if !(indexes.OptionsPerRelativeKey[indexEntry][0].Value == "RootLogin yes" && indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line == 1) {
		t.Errorf("Expected 'RootLogin yes', but got %v", indexes.OptionsPerRelativeKey[indexEntry][0].Value)
	}
}

func TestComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Port 22
Port 2022
Port 2024

Match Address 192.168.0.1/24
	PermitRootLogin no
	RoomLogin yes
	PermitRootLogin yes
`)
	config := ast.NewSSHConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if !(len(errors) == 1) {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexEntry := SSHIndexKey{
		Option:     "PermitRootLogin",
		MatchBlock: nil,
	}
	if !(indexes.OptionsPerRelativeKey[indexEntry][0].Value == "PermitRootLogin yes" && indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line == 0) {
		t.Errorf("Expected 'PermitRootLogin yes' on line 0, but got %v on line %v", indexes.OptionsPerRelativeKey[indexEntry][0].Value, indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line)
	}

	firstMatchBlock := config.FindMatchBlock(uint32(6))
	indexEntry = SSHIndexKey{
		Option:     "PermitRootLogin",
		MatchBlock: firstMatchBlock,
	}
	if !(indexes.OptionsPerRelativeKey[indexEntry][0].Value == "\tPermitRootLogin no" && indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line == 6) {
		t.Errorf("Expected 'PermitRootLogin no' on line 6, but got %v on line %v", indexes.OptionsPerRelativeKey[indexEntry][0].Value, indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line)
	}

	// Double check
	indexEntry = SSHIndexKey{
		Option:     "Port",
		MatchBlock: nil,
	}
	if !(indexes.OptionsPerRelativeKey[indexEntry][0].Value == "Port 22" &&
		indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line == 1 &&
		len(indexes.OptionsPerRelativeKey[indexEntry]) == 3 &&
		indexes.OptionsPerRelativeKey[indexEntry][1].Value == "Port 2022" &&
		indexes.OptionsPerRelativeKey[indexEntry][1].Start.Line == 2 &&
		indexes.OptionsPerRelativeKey[indexEntry][2].Value == "Port 2024" &&
		indexes.OptionsPerRelativeKey[indexEntry][2].Start.Line == 3) {
		t.Errorf("Expected 'Port 22' on line 1, but got %v on line %v", indexes.OptionsPerRelativeKey[indexEntry][0].Value, indexes.OptionsPerRelativeKey[indexEntry][0].Start.Line)
	}
}
