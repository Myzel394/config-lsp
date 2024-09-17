package indexes

import (
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/utils"
	"testing"
)

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
		t.Fatalf("Expected one errors, but got %v", len(errors))
	}

	firstMatchBlock := config.FindMatchBlock(uint32(6))
	opts := indexes.AllOptionsPerName["PermitRootLogin"]
	if !(len(opts) == 2 &&
		len(opts[nil]) == 1 &&
		opts[nil][0].Value.Value == "PermitRootLogin yes" &&
		opts[nil][0].Start.Line == 0 &&
		len(opts[firstMatchBlock]) == 1 &&
		opts[firstMatchBlock][0].Value.Value == "\tPermitRootLogin no" &&
		opts[firstMatchBlock][0].Start.Line == 6 &&
		opts[firstMatchBlock][0].Key.Key == "PermitRootLogin") {
		t.Errorf("Expected 3 PermitRootLogin options, but got %v", opts)
	}
}

func TestIncludeExample(
	t *testing.T,
) {
	input := utils.Dedent(`
PermitRootLogin yes
Include /etc/ssh/sshd_config.d/*.conf hello_world
`)
	config := ast.NewSSHConfig()
	errors := config.Parse(input)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	indexes, errors := CreateIndexes(*config)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", len(errors))
	}

	if !(len(indexes.Includes) == 1) {
		t.Fatalf("Expected 1 include, but got %v", len(indexes.Includes))
	}

	if !(len(indexes.Includes[1].Values) == 2) {
		t.Fatalf("Expected 2 include path, but got %v", len(indexes.Includes[1].Values))
	}

	if !(indexes.Includes[1].Values[0].Value == "/etc/ssh/sshd_config.d/*.conf" &&
		indexes.Includes[1].Values[0].Start.Line == 1 &&
		indexes.Includes[1].Values[0].End.Line == 1 &&
		indexes.Includes[1].Values[0].Start.Character == 8 &&
		indexes.Includes[1].Values[0].End.Character == 36) {
		t.Errorf("Expected '/etc/ssh/sshd_config.d/*.conf' on line 1, but got %v on line %v", indexes.Includes[1].Values[0].Value, indexes.Includes[1].Values[0].Start.Line)
	}

	if !(indexes.Includes[1].Values[1].Value == "hello_world" &&
		indexes.Includes[1].Values[1].Start.Line == 1 &&
		indexes.Includes[1].Values[1].End.Line == 1 &&
		indexes.Includes[1].Values[1].Start.Character == 38 &&
		indexes.Includes[1].Values[1].End.Character == 48) {
		t.Errorf("Expected 'hello_world' on line 1, but got %v on line %v", indexes.Includes[1].Values[1].Value, indexes.Includes[1].Values[1].Start.Line)
	}
}
