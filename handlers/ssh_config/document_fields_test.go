package sshconfig

import (
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/indexes"
	"config-lsp/utils"
	"testing"
)

func TestComplexExample(
	t *testing.T,
) {
	input := utils.Dedent(`
ProxyCommand hello

Host laptop
    HostName laptop.lan
	ProxyCommand test

Match originalhost laptop exec "[[ $(/usr/bin/dig +short laptop.lan) == '' ]]"
    HostName laptop.sdn
`)
	c := ast.NewSSHConfig()
	errors := c.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	i, errors := indexes.CreateIndexes(*c)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	d := &SSHDocument{
		Config:  c,
		Indexes: i,
	}

	options := d.FindOptionsByName("ProxyCommand")
	if !(len(options) == 2 && options[0].Option.Start.Line == 0 && options[1].Option.Start.Line == 4) {
		t.Errorf("Expected 2 options, got %v", options)
	}

	options = d.FindOptionsByName("HostName")
	if !(len(options) == 2 && options[0].Option.Start.Line == 3 && options[1].Option.Start.Line == 7) {
		t.Errorf("Expected 2 options, got %v", options)
	}

	block := d.Config.FindBlock(4)
	if !(d.FindOptionByNameAndBlock("ProxyCommand", block).Option.Start.Line == 4) {
		t.Errorf("Expected 4, got %v", d.FindOptionByNameAndBlock("PorxyCommand", block).Option.Start.Line)
	}

	if !(d.FindOptionByNameAndBlock("ProxyCommand", nil).Option.Start.Line == 0) {
		t.Errorf("Expected 0, got %v", d.FindOptionByNameAndBlock("ProxyCommand", nil).Option.Start.Line)
	}

	matchBlocks := d.GetAllMatchBlocks()
	if !(len(matchBlocks) == 1 && matchBlocks[0].Start.Line == 6) {
		t.Errorf("Expected 1 match block, got %v", matchBlocks)
	}
}
