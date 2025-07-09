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
}
