package matchparser

import (
	"testing"
)

func TestFullExample(
	t *testing.T,
) {
	input := "User alice,adam Address 192.168.1.1 Host *.example.com"
	match := NewMatch()
	errs := match.Parse(input, 0, 0)

	if len(errs) > 0 {
		t.Fatalf("Failed to parse match: %v", errs)
	}

	entry := match.GetEntryByCursor(0)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 0 to be %v, but got %v", match.Entries[0], entry)
	}

	entry = match.GetEntryByCursor(5)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 5 to be %v, but got %v", match.Entries[0], entry)
	}

	entry = match.GetEntryByCursor(13)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 13 to be %v, but got %v", match.Entries[1], entry)
	}

	entry = match.GetEntryByCursor(16)
	if !(entry == match.Entries[1]) {
		t.Errorf("Expected entry at 16 to be %v, but got %v", match.Entries[1], entry)
	}

	entry = match.GetEntryByCursor(24)
	if !(entry == match.Entries[1]) {
		t.Errorf("Expected entry at 24 to be %v, but got %v", match.Entries[2], entry)
	}

	entry = match.GetEntryByCursor(36)
	if !(entry == match.Entries[2]) {
		t.Errorf("Expected entry at 36 to be %v, but got %v", match.Entries[2], entry)
	}
}

func TestGetEntryForIncompleteExample(
	t *testing.T,
) {
	input := "User "
	match := NewMatch()
	errs := match.Parse(input, 0, 0)

	if len(errs) > 0 {
		t.Fatalf("Failed to parse match: %v", errs)
	}

	entry := match.GetEntryByCursor(0)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 0 to be %v, but got %v", match.Entries[0], entry)
	}

	entry = match.GetEntryByCursor(4)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 4 to be %v, but got %v", match.Entries[0], entry)
	}

	entry = match.GetEntryByCursor(5)
	if !(entry == match.Entries[0]) {
		t.Errorf("Expected entry at 5 to be %v, but got %v", match.Entries[0], entry)
	}
}
