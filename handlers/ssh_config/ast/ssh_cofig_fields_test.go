package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestComplexExampleRetrievesCorrectly(
	t *testing.T,
) {
	input := utils.Dedent(`
Port 22

Host laptop
    HostName laptop.lan

Match originalhost laptop exec "[[ $(/usr/bin/dig +short laptop.lan) == '' ]]"
    HostName laptop.sdn
	`)

	p := NewSSHConfig()
	errors := p.Parse(input)

	if len(errors) != 0 {
		t.Fatalf("Expected no errors, got %v", errors)
	}

	firstOption, firstBlock := p.FindOption(0)
	if !(firstOption.Value.Raw == "Port 22") {
		t.Errorf("Expected Port 22, got %v", firstOption.Value.Raw)
	}

	if !(firstBlock == nil) {
		t.Errorf("Expected no block, got %v", firstBlock)
	}

	secondOption, secondBlock := p.FindOption(3)
	if !(secondOption.Value.Raw == "    HostName laptop.lan") {
		t.Errorf("Expected HostName laptop.lan, got %v", secondOption.Value.Raw)
	}

	if !(secondBlock.GetLocation().Start.Line == 2) {
		t.Errorf("Expected line 2, got %v", secondBlock.GetLocation().Start.Line)
	}

	thirdOption, thirdBlock := p.FindOption(6)
	if !(thirdOption.Value.Raw == "    HostName laptop.sdn") {
		t.Errorf("Expected HostName laptop.sdn, got %v", thirdOption.Value.Raw)
	}

	if !(thirdBlock.GetLocation().Start.Line == 5) {
		t.Errorf("Expected line 3, got %v", thirdBlock.GetLocation().Start.Line)
	}
}
