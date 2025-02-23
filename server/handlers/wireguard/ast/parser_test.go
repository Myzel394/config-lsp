package ast

import (
	"config-lsp/utils"
	"testing"
)

func TestExample1Works(
	t *testing.T,
) {
	sample := utils.Dedent(`
# A comment at the very top


[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	config := NewWGConfig()

	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(utils.KeyExists(config.CommentLines, 0) && utils.KeyExists(config.CommentLines, 12)) {
		t.Errorf("Parse: Expected comments to be present on lines 0 and 12")
	}

	if !(config.Sections[0].Header.Name == "Interface" && config.Sections[1].Header.Name == "Peer") {
		t.Errorf("Parse: Expected sections to be present on lines 0, 1, and 2")
	}

	if !(config.Sections[0].Properties[4].Key.Name == "PrivateKey" && config.Sections[0].Properties[4].Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 4 to be correct")
	}

	if !(config.Sections[0].Properties[5].Key.Name == "Address" && config.Sections[0].Properties[5].Value.Value == "10.0.0.1") {
		t.Errorf("Parse: Expected property line 5 to be correct")
	}

	if !(config.Sections[1].Properties[10].Key.Name == "PublicKey" && config.Sections[1].Properties[10].Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 10 to be correct")
	}
}
