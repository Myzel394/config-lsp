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

// ========== Test negation ==========

func TestSimpleIPTableRuleNegation(t *testing.T) {
	rule := `iptables -I FORWARD -i wg0 -j ACCEPT`

	parsedRule, err := ParseIpTableRule(rule)

	if err != nil {
		t.Fatalf("Failed to parse rule: %v", err)
	}

	negatedRule := parsedRule.InvertAction().String()

	if negatedRule != `iptables -D FORWARD -i wg0 -j ACCEPT` {
		t.Errorf("Expected negated rule to be 'iptables -D FORWARD -i wg0 -j ACCEPT', got '%s'", negatedRule)
	}
}

func TestComplexIPTableRuleNegation(t *testing.T) {
	rule := `iptables -I FORWARD -i wg0 -o eth0 -p tcp --dport 80 -j ACCEPT`

	parsedRule, err := ParseIpTableRule(rule)

	if err != nil {
		t.Fatalf("Failed to parse rule: %v", err)
	}

	negatedRule := parsedRule.InvertAction().String()

	if negatedRule != `iptables -D FORWARD -i wg0 -o eth0 -p tcp --dport 80 -j ACCEPT` {
		t.Errorf("Expected negated rule to be 'iptables -D FORWARD -i wg0 -o eth0 -p tcp --dport 80 -j ACCEPT', got '%s'", negatedRule)
	}
}
