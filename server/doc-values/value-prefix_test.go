package docvalues

import "testing"

func TestPrefixValidExample(t *testing.T) {
	value := PrefixValue{
		Prefixes: []Prefix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		Required: true,
		SubValue: StringValue{},
	}

	errs := value.DeprecatedCheckIsValid("debug_example")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestPrefixInvalidExample(t *testing.T) {
	value := PrefixValue{
		Prefixes: []Prefix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		Required: true,
		SubValue: StringValue{},
	}

	errs := value.DeprecatedCheckIsValid("example")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestPrefixNotRequiredExample(t *testing.T) {
	value := PrefixValue{
		Prefixes: []Prefix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
		Required: false,
	}

	errs := value.DeprecatedCheckIsValid("example")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestPrefixNotRequiredButContainsPrefixExample(t *testing.T) {
	value := PrefixValue{
		Prefixes: []Prefix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
		Required: false,
	}

	errs := value.DeprecatedCheckIsValid("debug_example")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}
