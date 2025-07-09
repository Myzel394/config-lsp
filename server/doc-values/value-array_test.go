package docvalues

import (
	"regexp"
	"testing"
)

func TestVAValidExample1(t *testing.T) {
	value := ArrayValue{
		SubValue: RegexValue{
			Regex: *regexp.MustCompile(`^\d+$`),
		},
		Separator:           ",",
		RespectQuotes:       false,
		DuplicatesExtractor: nil,
	}

	line := "1,212,3445"
	errs := value.DeprecatedCheckIsValid(line)

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, val1Cursor := value.getCurrentValue(line, 0)
	if !(val1 == "1" && val1Cursor == 0) {
		t.Errorf("Expected first value to be '1' at cursor 0, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 1)
	if !(val1 == "1" && val1Cursor == 1) {
		t.Errorf("Expected first value to be '1' at cursor 1, got '%s' at %d", val1, val1Cursor)
	}

	val2, val2Cursor := value.getCurrentValue(line, 2)
	if !(val2 == "212" && val2Cursor == 0) {
		t.Errorf("Expected second value to be '212' at cursor 2, got '%s' at %d", val2, val2Cursor)
	}

	val2, val2Cursor = value.getCurrentValue(line, 3)
	if !(val2 == "212" && val2Cursor == 1) {
		t.Errorf("Expected second value to be '212' at cursor 3, got '%s' at %d", val2, val2Cursor)
	}

	val2, val2Cursor = value.getCurrentValue(line, 5)
	if !(val2 == "212" && val2Cursor == 3) {
		t.Errorf("Expected second value to be '212' at cursor 5, got '%s' at %d", val2, val2Cursor)
	}

	val3, val3Cursor := value.getCurrentValue(line, 6)
	if !(val3 == "3445" && val3Cursor == 0) {
		t.Errorf("Expected third value to be '3445' at cursor 6, got '%s' at %d", val3, val3Cursor)
	}

	val3, val3Cursor = value.getCurrentValue(line, 9)
	if !(val3 == "3445" && val3Cursor == 3) {
		t.Errorf("Expected third value to be '3445' at cursor 9, got '%s' at %d", val3, val3Cursor)
	}

	val3, val3Cursor = value.getCurrentValue(line, 10)
	if !(val3 == "3445" && val3Cursor == 4) {
		t.Errorf("Expected third value to be '3445' at cursor 10, got '%s' at %d", val3, val3Cursor)
	}
}

func TestVAEdgeCaseAtStart(t *testing.T) {
	value := ArrayValue{
		SubValue: RegexValue{
			Regex: *regexp.MustCompile(`^\d*$`),
		},
		Separator:           ",",
		RespectQuotes:       false,
		DuplicatesExtractor: nil,
	}

	line := ",1,212,3445"

	errs := value.DeprecatedCheckIsValid(line)

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, val1Cursor := value.getCurrentValue(line, 0)
	if !(val1 == "" && val1Cursor == 0) {
		t.Errorf("Expected first value to be '' at cursor 0, got '%s' at %d", val1, val1Cursor)
	}

	val2, val2Cursor := value.getCurrentValue(line, 1)
	if !(val2 == "1" && val2Cursor == 0) {
		t.Errorf("Expected second value to be '1' at cursor 1, got '%s' at %d", val2, val2Cursor)
	}

	val2, val2Cursor = value.getCurrentValue(line, 2)
	if !(val2 == "1" && val2Cursor == 1) {
		t.Errorf("Expected second value to be '1' at cursor 2, got '%s' at %d", val2, val2Cursor)
	}

	val3, val3Cursor := value.getCurrentValue(line, 3)
	if !(val3 == "212" && val3Cursor == 0) {
		t.Errorf("Expected third value to be '212' at cursor 3, got '%s' at %d", val3, val3Cursor)
	}
}

func TestVAEdgeCaseAtMiddle(t *testing.T) {
	value := ArrayValue{
		SubValue: RegexValue{
			Regex: *regexp.MustCompile(`^\d*$`),
		},
		Separator:           ",",
		RespectQuotes:       false,
		DuplicatesExtractor: nil,
	}

	line := "1,,212,3445"

	errs := value.DeprecatedCheckIsValid(line)

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, val1Cursor := value.getCurrentValue(line, 0)
	if !(val1 == "1" && val1Cursor == 0) {
		t.Errorf("Expected first value to be '1' at cursor 0, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 1)
	if !(val1 == "1" && val1Cursor == 1) {
		t.Errorf("Expected first value to be '1' at cursor 1, got '%s' at %d", val1, val1Cursor)
	}

	val2, val2Cursor := value.getCurrentValue(line, 2)
	if !(val2 == "" && val2Cursor == 0) {
		t.Errorf("Expected second value to be '' at cursor 2, got '%s' at %d", val2, val2Cursor)
	}

	val3, val3Cursor := value.getCurrentValue(line, 3)
	if !(val3 == "212" && val3Cursor == 0) {
		t.Errorf("Expected third value to be '212' at cursor 3, got '%s' at %d", val3, val3Cursor)
	}
}

func TestVAEdgeCaseAtEnd(t *testing.T) {
	value := ArrayValue{
		SubValue: RegexValue{
			Regex: *regexp.MustCompile(`^\d*$`),
		},
		Separator:           ",",
		RespectQuotes:       false,
		DuplicatesExtractor: nil,
	}

	line := "1,"

	errs := value.DeprecatedCheckIsValid(line)

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, val1Cursor := value.getCurrentValue(line, 0)
	if !(val1 == "1" && val1Cursor == 0) {
		t.Errorf("Expected first value to be '1' at cursor 0, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 1)
	if !(val1 == "1" && val1Cursor == 1) {
		t.Errorf("Expected first value to be '1' at cursor 1, got '%s' at %d", val1, val1Cursor)
	}

	val2, val2Cursor := value.getCurrentValue(line, 2)
	if !(val2 == "" && val2Cursor == 0) {
		t.Errorf("Expected second value to be '' at cursor 2, got '%s' at %d", val2, val2Cursor)
	}
}

func TestVAValidExampleQuotes(t *testing.T) {
	value := ArrayValue{
		SubValue:      StringValue{},
		Separator:     ",",
		RespectQuotes: true,
		PersistQuotes: false,
	}

	line := `"1,212",3445`

	errs := value.DeprecatedCheckIsValid(line)

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, val1Cursor := value.getCurrentValue(line, 0)
	if !(val1 == `"1,212"` && val1Cursor == 0) {
		t.Errorf("Expected first value to be '1,212' at cursor 0, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 1)
	if !(val1 == `"1,212"` && val1Cursor == 1) {
		t.Errorf("Expected first value to be '1,212' at cursor 1, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 2)
	if !(val1 == `"1,212"` && val1Cursor == 2) {
		t.Errorf("Expected first value to be '1,212' at cursor 2, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 6)
	if !(val1 == `"1,212"` && val1Cursor == 6) {
		t.Errorf("Expected first value to be '1,212' at cursor 5, got '%s' at %d", val1, val1Cursor)
	}

	val1, val1Cursor = value.getCurrentValue(line, 7)
	if !(val1 == `"1,212"` && val1Cursor == 7) {
		t.Errorf("Expected first value to be '1,212' at cursor 6, got '%s' at %d", val1, val1Cursor)
	}

	val2, val2Cursor := value.getCurrentValue(line, 8)
	if !(val2 == "3445" && val2Cursor == 0) {
		t.Errorf("Expected second value to be '3445' at cursor 8, got '%s' at %d", val2, val2Cursor)
	}

	val2, val2Cursor = value.getCurrentValue(line, 9)
	if !(val2 == "3445" && val2Cursor == 1) {
		t.Errorf("Expected second value to be '3445' at cursor 9, got '%s' at %d", val2, val2Cursor)
	}

	val2, val2Cursor = value.getCurrentValue(line, 12)
	if !(val2 == "3445" && val2Cursor == 4) {
		t.Errorf("Expected second value to be '3445' at cursor 12, got '%s' at %d", val2, val2Cursor)
	}
}
