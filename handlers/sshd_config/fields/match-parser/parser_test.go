package match_parser

import (
	"testing"
)

func TestComplexExample(
	t *testing.T,
) {
	input := "User root,admin,alice Address *,!192.168.0.1"

	match := NewMatch()
	errors := match.Parse(input, 32)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 2) {
		t.Fatalf("Expected 2 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria == MatchCriteriaTypeUser) {
		t.Fatalf("Expected User, but got %v", match.Entries[0])
	}

	if !(match.Entries[0].Values[0].Value == "root") {
		t.Fatalf("Expected root, but got %v", match.Entries[0].Values[0])
	}

	if !(match.Entries[0].Values[1].Value == "admin") {
		t.Fatalf("Expected admin, but got %v", match.Entries[0].Values[1])
	}

	if !(match.Entries[0].Values[2].Value == "alice") {
		t.Fatalf("Expected alice, but got %v", match.Entries[0].Values[2])
	}

	if !(match.Entries[1].Criteria == MatchCriteriaTypeAddress) {
		t.Fatalf("Expected Address, but got %v", match.Entries[1])
	}

	if !(match.Entries[1].Values[0].Value == "*") {
		t.Fatalf("Expected *, but got %v", match.Entries[1].Values[0])
	}

	if !(match.Entries[1].Values[1].Value == "!192.168.0.1") {
		t.Fatalf("Expected !192.168.0.1, but got %v", match.Entries[1].Values[1])
	}
}

func TestSecondComplexExample(
	t *testing.T,
) {
	input := "Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1"

	match := NewMatch()
	errors := match.Parse(input, 0)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 1) {
		t.Fatalf("Expected 1 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria == MatchCriteriaTypeAddress) {
		t.Fatalf("Expected Address, but got %v", match.Entries[0])
	}

	if !(len(match.Entries[0].Values) == 3) {
		t.Fatalf("Expected 3 values, but got %v", len(match.Entries[0].Values))
	}

	if !(match.Entries[0].Values[0].Value == "172.22.100.0/24") {
		t.Fatalf("Expected 172.22.100.0/24, but got %v", match.Entries[0].Values[0])
	}
}
