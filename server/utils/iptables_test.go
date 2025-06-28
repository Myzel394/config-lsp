package utils

import "testing"

func TestSimpleIPTableRuleParsing(t *testing.T) {
	rule := `iptables -I FORWARD -i wg0 -j ACCEPT`

	parsedRule, err := ParseIpTableRule(rule)

	if err != nil {
		t.Fatalf("Failed to parse rule: %v", err)
	}

	if parsedRule.Action != IpTableActionInsert {
		t.Errorf("Expected action to be %v, got %v", IpTableActionInsert, parsedRule.Action)
	}

	if parsedRule.Command != "iptables  FORWARD -i wg0 -j ACCEPT" {
		t.Errorf("Expected command to be 'iptables  FORWARD -i wg0 -j ACCEPT', got '%s'", parsedRule.Command)
	}

	if parsedRule.ActionIndex != 9 {
		t.Errorf("Expected action index to be 9, got %d", parsedRule.ActionIndex)
	}
}

func TestSimpleIPTableRule2Parsing(t *testing.T) {
	rule := `iptables --insert FORWARD -i wg0 -j ACCEPT`

	_, err := ParseIpTableRule(rule)

	if err != nil {
		t.Fatalf("Failed to parse rule: %v", err)
	}
}

func TestInvalidIPTableRuleParsing(t *testing.T) {
	rule := `iptables -i wg0 -j ACCEPT`

	_, err := ParseIpTableRule(rule)

	if err == nil {
		t.Fatal("Expected error for invalid rule, but got none")
	}
}
