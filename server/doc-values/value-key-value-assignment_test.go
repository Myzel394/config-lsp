package docvalues

import "testing"

func TestKVAValidExample1(t *testing.T) {
	value := KeyValueAssignmentValue{
		Key:       StringValue{},
		Value:     NumberValue{},
		Separator: "=",
	}

	errs := value.DeprecatedCheckIsValid("key=123")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}

	val1, select1, cursor1 := value.getValueAtCursor("key=123", 0)
	if !(val1 == "key" && *select1 == keySelected && cursor1 == 0) {
		t.Errorf("Expected (key, keySelected, 0), got (%s, %v, %d)", val1, select1, cursor1)
	}

	val2, select2, cursor2 := value.getValueAtCursor("key=123", 2)
	if !(val2 == "key" && *select2 == keySelected && cursor2 == 2) {
		t.Errorf("Expected (key=123, valueSelected, 2), got (%s, %v, %d)", val2, select2, cursor2)
	}

	val3, select3, cursor3 := value.getValueAtCursor("key=123", 4)
	if !(val3 == "123" && *select3 == valueSelected && cursor3 == 0) {
		t.Errorf("Expected (key=123, valueSelected, 0), got (%s, %v, %d)", val3, select3, cursor3)
	}

	val4, select4, cursor4 := value.getValueAtCursor("key=123", 6)
	if !(val4 == "123" && *select4 == valueSelected && cursor4 == 2) {
		t.Errorf("Expected (key=123, valueSelected, 2), got (%s, %v, %d)", val4, select4, cursor4)
	}

	val5, select5, cursor5 := value.getValueAtCursor("key=123", 3)
	if !(val5 == "" && select5 == nil && cursor5 == 0) {
		t.Errorf("Expected (key=123, nil, 0), got (%s, %v, %d)", val5, select5, cursor5)
	}
}
