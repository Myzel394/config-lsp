package common

import (
	"testing"
)

func TestCursorPosition(
	t *testing.T,
) {
	// Contains fictive range for the name "Test" in the code:
	// func Test() {}
	locationRange := LocationRange{
		Start: Location{
			Line:      0,
			Character: 5,
		},
		End: Location{
			Line:      0,
			Character: 9,
		},
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsCursorPosition(5)) == true) {
		t.Errorf("Expected 5 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsCursorPosition(6)) == true) {
		t.Errorf("Expected 6 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsCursorPosition(9)) == true) {
		t.Errorf("Expected 9 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsCursorPosition(10)) == false) {
		t.Errorf("Expected 10 to not be in range, but it was")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsCursorPosition(4)) == false) {
		t.Errorf("Expected 4 to not be in range, but it was")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsCursorPosition(0)) == true) {
		t.Errorf("Expected 0 to be before start, but it wasn't")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsCursorPosition(4)) == true) {
		t.Errorf("Expected 5 to be before start, but it wasn't")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsCursorPosition(5)) == false) {
		t.Errorf("Expected 5 to not be before start, but it was")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsCursorPosition(10)) == false) {
		t.Errorf("Expected 10 to not be before start, but it was")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsCursorPosition(10)) == true) {
		t.Errorf("Expected 10 to be after end, but it wasn't")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsCursorPosition(11)) == true) {
		t.Errorf("Expected 11 to be after end, but it wasn't")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsCursorPosition(9)) == false) {
		t.Errorf("Expected 9 to not be after end, but it was")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsCursorPosition(5)) == false) {
		t.Errorf("Expected 5 to not be after end, but it was")
	}
}

func TestIndexPosition(t *testing.T) {
	// Contains fictive range for the name "Test" in the code:
	// func Test() {}
	locationRange := LocationRange{
		Start: Location{
			Line:      0,
			Character: 5,
		},
		End: Location{
			Line:      0,
			Character: 9,
		},
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(5)) == true) {
		t.Errorf("Expected index position 5 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(6)) == true) {
		t.Errorf("Expected index position 6 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(8)) == true) {
		t.Errorf("Expected index position 6 to be in range, but it wasn't")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(9)) == false) {
		t.Errorf("Expected index position 9 to not be in range, but it was")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(10)) == false) {
		t.Errorf("Expected index position 10 to not be in range, but it was")
	}

	if !(locationRange.ContainsPosition(LSPCharacterAsIndexPosition(4)) == false) {
		t.Errorf("Expected index position 4 to not be in range, but it was")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsIndexPosition(4)) == true) {
		t.Errorf("Expected index position 4 to be before start, but it wasn't")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsIndexPosition(5)) == false) {
		t.Errorf("Expected index position 5 to not be before start, but it was")
	}

	if !(locationRange.IsPositionBeforeStart(LSPCharacterAsIndexPosition(10)) == false) {
		t.Errorf("Expected index position 10 to not be before start, but it wasn't")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsIndexPosition(10)) == true) {
		t.Errorf("Expected index position 10 to be after end, but it wasn't")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsIndexPosition(9)) == true) {
		t.Errorf("Expected index position 9 to be after end, but it wasn't")
	}

	if !(locationRange.IsPositionAfterEnd(LSPCharacterAsIndexPosition(5)) == false) {
		t.Errorf("Expected index position 5 to not be after end, but it was")
	}
}
