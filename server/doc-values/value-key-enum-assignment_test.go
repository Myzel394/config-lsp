package docvalues

import "testing"

func TestKEAValidExample1(t *testing.T) {
	value := KeyEnumAssignmentValue{
		Values: map[EnumString]DeprecatedValue{
			CreateEnumString("burger"): NumberValue{},
			CreateEnumString("pizza"):  StringValue{},
		},
		Separator: "=",
	}

	errs := value.DeprecatedCheckIsValid("burger=12")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestKEAValidExampleValueOptional(t *testing.T) {
	value := KeyEnumAssignmentValue{
		Values: map[EnumString]DeprecatedValue{
			CreateEnumString("burger"): NumberValue{},
			CreateEnumString("pizza"):  StringValue{},
		},
		Separator:       "=",
		ValueIsOptional: true,
	}

	errs := value.DeprecatedCheckIsValid("burger")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestKEAValidExample2(t *testing.T) {
	value := KeyEnumAssignmentValue{
		Values: map[EnumString]DeprecatedValue{
			CreateEnumString("burger"): NumberValue{},
			CreateEnumString("pizza"):  StringValue{},
		},
		Separator: "=",
	}

	errs := value.DeprecatedCheckIsValid("pizza=cheese")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestKEAInvalidExample1(t *testing.T) {
	value := KeyEnumAssignmentValue{
		Values: map[EnumString]DeprecatedValue{
			CreateEnumString("burger"): NumberValue{},
			CreateEnumString("pizza"):  StringValue{},
		},
		Separator: "=",
	}

	errs := value.DeprecatedCheckIsValid("burger=cheese")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	} else {
		t.Logf("Got expected errors: %v", errs)
	}
}

func TestKEAInvalidExampleEnumDoesNotExist(t *testing.T) {
	value := KeyEnumAssignmentValue{
		Values: map[EnumString]DeprecatedValue{
			CreateEnumString("burger"): NumberValue{},
			CreateEnumString("pizza"):  StringValue{},
		},
		Separator: "=",
	}

	errs := value.DeprecatedCheckIsValid("salad=12")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	} else {
		t.Logf("Got expected errors: %v", errs)
	}
}
