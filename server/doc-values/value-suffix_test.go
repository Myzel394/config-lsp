package docvalues

import "testing"

func TestSuffixValidExample(t *testing.T) {
	value := SuffixValue{
		Suffixes: []Suffix{
			{Suffix: "_debug", Meaning: "Enable debug mode"},
			{Suffix: "_test", Meaning: "Enable test mode"},
		},
		Required: true,
		SubValue: StringValue{},
	}

	errs := value.DeprecatedCheckIsValid("example_debug")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestSuffixInvalidExample(t *testing.T) {
	value := SuffixValue{
		Suffixes: []Suffix{
			{Suffix: "_debug", Meaning: "Enable debug mode"},
			{Suffix: "_test", Meaning: "Enable test mode"},
		},
		Required: true,
		SubValue: StringValue{},
	}

	errs := value.DeprecatedCheckIsValid("example")

	if len(errs) == 0 {
		t.Error("Expected errors, got none")
	}
}

func TestSuffixNotRequiredExample(t *testing.T) {
	value := SuffixValue{
		Suffixes: []Suffix{
			{Suffix: "_debug", Meaning: "Enable debug mode"},
			{Suffix: "_test", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
		Required: false,
	}

	errs := value.DeprecatedCheckIsValid("example")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}

func TestSuffixNotRequiredButContainsSuffixExample(t *testing.T) {
	value := SuffixValue{
		Suffixes: []Suffix{
			{Suffix: "_debug", Meaning: "Enable debug mode"},
			{Suffix: "_test", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
		Required: false,
	}

	errs := value.DeprecatedCheckIsValid("example_debug")

	if len(errs) != 0 {
		t.Errorf("Expected no errors, got: %v", errs)
	}
}
