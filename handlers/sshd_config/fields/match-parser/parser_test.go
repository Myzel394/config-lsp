package match_parser

import (
	"testing"
)

func TestComplexExample(
	t *testing.T,
) {
	offset := uint32(5)
	input := "User root,admin,alice Address *,!192.168.0.1"

	match := NewMatch()
	errors := match.Parse(input, 32, offset)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 2) {
		t.Fatalf("Expected 2 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeUser) {
		t.Fatalf("Expected User, but got %v", match.Entries[0])
	}

	if !(match.Entries[0].Values.Values[0].Value.Value == "root" && match.Entries[0].Values.Values[0].Start.Character == 5+offset && match.Entries[0].Values.Values[0].End.Character == 8+offset && match.Entries[0].Start.Character == 0+offset && match.Entries[0].End.Character == 20+offset) {
		t.Errorf("Expected root, but got %v", match.Entries[0].Values.Values[0])
	}

	if !(match.Entries[0].Values.Values[1].Value.Value == "admin" && match.Entries[0].Values.Values[1].Start.Character == 10+offset && match.Entries[0].Values.Values[1].End.Character == 14+offset) {
		t.Errorf("Expected admin, but got %v", match.Entries[0].Values.Values[1])
	}

	if !(match.Entries[0].Values.Values[2].Value.Value == "alice" && match.Entries[0].Values.Values[2].Start.Character == 16+offset && match.Entries[0].Values.Values[2].End.Character == 20+offset) {
		t.Errorf("Expected alice, but got %v", match.Entries[0].Values.Values[2])
	}

	if !(match.Entries[1].Criteria.Type == MatchCriteriaTypeAddress) {
		t.Errorf("Expected Address, but got %v", match.Entries[1])
	}

	if !(match.Entries[1].Values.Values[0].Value.Value == "*" && match.Entries[1].Values.Values[0].Start.Character == 30+offset && match.Entries[1].Values.Values[0].End.Character == 30+offset) {
		t.Errorf("Expected *, but got %v", match.Entries[1].Values.Values[0])
	}

	if !(match.Entries[1].Values.Values[1].Value.Value == "!192.168.0.1" && match.Entries[1].Values.Values[1].Start.Character == 32+offset && match.Entries[1].Values.Values[1].End.Character == 43+offset) {
		t.Errorf("Expected !192.168.0.1, but got %v", match.Entries[1].Values.Values[1])
	}
}

func TestSecondComplexExample(
	t *testing.T,
) {
	input := "Address 172.22.100.0/24,172.22.5.0/24,127.0.0.1"

	match := NewMatch()
	errors := match.Parse(input, 0, 20)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 1) {
		t.Fatalf("Expected 1 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeAddress) {
		t.Fatalf("Expected Address, but got %v", match.Entries[0])
	}

	if !(len(match.Entries[0].Values.Values) == 3) {
		t.Fatalf("Expected 3 values, but got %v", len(match.Entries[0].Values.Values))
	}

	if !(match.Entries[0].Values.Values[0].Value.Value == "172.22.100.0/24") {
		t.Fatalf("Expected 172.22.100.0/24, but got %v", match.Entries[0].Values.Values[0])
	}
}

func TestIncompleteBetweenEntriesExample(
	t *testing.T,
) {
	input := "User root,admin,alice "

	match := NewMatch()
	errors := match.Parse(input, 0, 0)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 1) {
		t.Errorf("Expected 1 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeUser) {
		t.Errorf("Expected User, but got %v", match.Entries[0])
	}

	if !(len(match.Entries[0].Values.Values) == 3) {
		t.Errorf("Expected 3 values, but got %v", len(match.Entries[0].Values.Values))
	}

	if !(match.Entries[0].Start.Character == 0 && match.Entries[0].End.Character == 20) {
		t.Errorf("Expected 0-20, but got %v", match.Entries[0])
	}
}

func TestIncompleteBetweenValuesExample(
	t *testing.T,
) {
	input := "User "

	match := NewMatch()
	errors := match.Parse(input, 0, 0)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 1) {
		t.Errorf("Expected 1 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeUser) {
		t.Errorf("Expected User, but got %v", match.Entries[0])
	}

	if !(match.Entries[0].Values == nil) {
		t.Errorf("Expected 0 values, but got %v", match.Entries[0].Values)
	}
}
