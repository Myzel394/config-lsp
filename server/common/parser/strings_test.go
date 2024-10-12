package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStringsSingleWortQuotedFullFeatures(
	t *testing.T,
) {
	input := `hello "world"`
	expected := ParsedString{
		Raw:   input,
		Value: "hello world",
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsFullyQuotedFullFeatures(
	t *testing.T,
) {
	input := `"hello world"`
	expected := ParsedString{
		Raw:   input,
		Value: "hello world",
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsMultipleQuotesFullFeatures(
	t *testing.T,
) {
	input := `hello "world goodbye"`
	expected := ParsedString{
		Raw:   input,
		Value: "hello world goodbye",
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTrimWhitespaceNoQuotes(
	t *testing.T,
) {
	input := "  hello    world  "
	expected := "hello world"

	actual := TrimWhitespace(input, false)

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestTrimWhitespaceQuotes(
	t *testing.T,
) {
	input := `  "hello    world"  `
	expected := `"hello    world"`

	actual := TrimWhitespace(input, true)

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsSimpleEscapedFullFeatures(
	t *testing.T,
) {
	input := `hello \"world`
	expected := ParsedString{
		Raw:   input,
		Value: `hello "world`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsEscapedQuotesFullFeatures(
	t *testing.T,
) {
	input := `hello \"world\"`
	expected := ParsedString{
		Raw:   input,
		Value: `hello "world"`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsQuotesAndEscapedFullFeatures(
	t *testing.T,
) {
	input := `hello "world how\" are you"`
	expected := ParsedString{
		Raw:   input,
		Value: `hello world how" are you`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsIncompleteQuotesFullFeatures(
	t *testing.T,
) {
	input := `hello "world`
	expected := ParsedString{
		Raw:   input,
		Value: `hello "world`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsIncompleteQuoteEscapedFullFeatures(
	t *testing.T,
) {
	input := `hello "world\"`
	expected := ParsedString{
		Raw:   input,
		Value: `hello "world"`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsIncompleteQuotes2FullFeatures(
	t *testing.T,
) {
	input := `hello "world how" "are you`
	expected := ParsedString{
		Raw:   input,
		Value: `hello world how "are you`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsIncompleteQuotes3FullFeatures(
	t *testing.T,
) {
	input := `hello "world how are you`
	expected := ParsedString{
		Raw:   input,
		Value: `hello "world how are you`,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsReplacements(
	t *testing.T,
) {
	input := `Hello\\040World`
	expected := ParsedString{
		Raw:   input,
		Value: `Hello World`,
	}

	actual := ParseRawString(input, ParseFeatures{
		ParseDoubleQuotes:      true,
		ParseEscapedCharacters: true,
		Replacements: &map[string]string{
			`\\040`: " ",
		},
	})

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
