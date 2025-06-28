package ini

import (
	"config-lsp/utils"
	"testing"
)

func TestParserWorks(t *testing.T) {
	sample := utils.Dedent(`
# A comment at the very top


[Interface]
PrivateKey = 1234567890 # Some comment
Address = 10.0.0.1
PostUp = iptables -I FORWARD -i wg0 -j ACCEPT; iptables -I INPUT -i wg0 -j ACCEPT



[Peer]
PublicKey = 1234567890

; I'm a comment
`)

	config := NewConfig()

	errors := config.Parse(sample)

	if len(errors) > 0 {
		t.Fatalf("Parse: Expected no errors, but got %v", errors)
	}

	if !(utils.KeyExists(config.CommentLines, 0) && utils.KeyExists(config.CommentLines, 13)) {
		t.Errorf("Parse: Expected comments to be present on lines 0 and 13")
	}

	if !(config.Sections[0].Start.Line == 3 && config.Sections[0].End.Line == 9) {
		t.Errorf("Parse: Expected section 0 to be present on lines 3 and 9, but it is: %v", config.Sections[0].End)
	}

	if !(config.Sections[0].Start.Character == 0 && config.Sections[0].End.Character == 0) {
		t.Errorf("Parse: Expected section 0 to be present on characters 0 and 0, but it is: %v", config.Sections[0].End)
	}

	if !(config.Sections[0].Header.Name == "Interface" && config.Sections[1].Header.Name == "Peer") {
		t.Errorf("Parse: Expected sections to be present on lines 0, 1, and 2")
	}

	rawFourthProperty, _ := config.Sections[0].Properties.Get(uint32(4))
	fourthProperty := rawFourthProperty.(*Property)
	if !(fourthProperty.Key.Name == "PrivateKey" && fourthProperty.Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 4 to be correct")
	}

	rawFifthProperty, _ := config.Sections[0].Properties.Get(uint32(5))
	fifthProperty := rawFifthProperty.(*Property)
	if !(fifthProperty.Key.Name == "Address" && fifthProperty.Value.Value == "10.0.0.1") {
		t.Errorf("Parse: Expected property line 5 to be correct")
	}

	rawTenthProperty, _ := config.Sections[1].Properties.Get(uint32(11))
	tenthProperty := rawTenthProperty.(*Property)
	if !(tenthProperty.Key.Name == "PublicKey" && tenthProperty.Value.Value == "1234567890") {
		t.Errorf("Parse: Expected property line 11 to be correct")
	}

	rawPostUpProperty, _ := config.Sections[0].Properties.Get(uint32(6))
	postUpProperty := rawPostUpProperty.(*Property)
	if !(postUpProperty.Key.Name == "PostUp" && postUpProperty.Value.Value == "iptables -I FORWARD -i wg0 -j ACCEPT; iptables -I INPUT -i wg0 -j ACCEPT") {
		t.Errorf("Parse: Expected PostUp property to be correct; %v; %v", postUpProperty, postUpProperty.Value.Value)
	}
}
