package matchparser

import (
	"testing"
)

func TestSimpleExample(
	t *testing.T,
) {
	offset := uint32(5)
	input := "originalhost example.com user root"

	match := NewMatch()
	errors := match.Parse(input, 32, offset)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 2) {
		t.Errorf("Expected 2 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeOriginalHost) {
		t.Errorf("Expected OriginalHost, but got %v", match.Entries[0])
	}

	if !(match.Entries[0].Values.Values[0].Value.Value == "example.com" && match.Entries[0].Values.Values[0].Start.Character == 13+offset && match.Entries[0].Values.Values[0].End.Character == 23+offset+1) {
		t.Errorf("Expected example.com, but got %v", match.Entries[0].Values.Values[0])
	}

	if !(match.Entries[1].Criteria.Type == MatchCriteriaTypeUser) {
		t.Errorf("Expected User, but got %v", match.Entries[1])
	}

	if !(match.Entries[1].Values.Values[0].Value.Value == "root" && match.Entries[1].Values.Values[0].Start.Character == 30+offset && match.Entries[1].Values.Values[0].End.Character == 33+offset+1) {
		t.Errorf("Expected root, but got %v", match.Entries[1].Values.Values[0])
	}
}

func TestListExample(
	t *testing.T,
) {
	offset := uint32(20)
	input := "originalhost example.com,example.org,example.net"

	match := NewMatch()
	errors := match.Parse(input, 0, offset)

	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 1) {
		t.Errorf("Expected 1 entries, but got %v", len(match.Entries))
	}

	if !(match.Entries[0].Criteria.Type == MatchCriteriaTypeOriginalHost) {
		t.Errorf("Expected Address, but got %v", match.Entries[0])
	}

	if !(len(match.Entries[0].Values.Values) == 3) {
		t.Errorf("Expected 3 values, but got %v", len(match.Entries[0].Values.Values))
	}

	if !(match.Entries[0].Values.Values[0].Value.Value == "example.com" && match.Entries[0].Values.Values[0].Start.Character == 13+offset && match.Entries[0].Values.Values[0].End.Character == 23+offset+1) {
		t.Errorf("Expected example.com, but got %v", match.Entries[0].Values.Values[0])
	}
}

func TestComplexExample(
	t *testing.T,
) {
	input := `originalhost laptop exec "[[ $(/usr/bin/dig +short laptop.lan) == '' ]]"`

	match := NewMatch()
	errors := match.Parse(input, 0, 0)

	// TODO: Fix match so that it allows quotes
	if len(errors) > 0 {
		t.Fatalf("Expected no errors, but got %v", errors)
	}

	if !(len(match.Entries) == 2) {
		t.Errorf("Expected 2 entries, but got %v", len(match.Entries))
	}
}

func TestIncompleteBetweenEntriesExample(
	t *testing.T,
) {
	input := "user root,admin,alice "

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

	if !(match.Entries[0].Start.Character == 0 && match.Entries[0].End.Character == 21) {
		t.Errorf("Expected 0-20, but got %v", match.Entries[0])
	}
}

func TestIncompleteBetweenValuesExample(
	t *testing.T,
) {
	input := "user "

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

	if !(len(match.Entries[0].Values.Values) == 0) {
		t.Errorf("Expected 0 values, but got %v", match.Entries[0].Values)
	}
}


