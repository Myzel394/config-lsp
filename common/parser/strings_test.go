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
		Raw:      input,
		Value:    "hello world",
		Features: FullFeatures,
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
		Raw:      input,
		Value:    "hello world",
		Features: FullFeatures,
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
		Raw:      input,
		Value:    "hello world goodbye",
		Features: FullFeatures,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringsSimpleEscapedFullFeatures(
	t *testing.T,
) {
	input := `hello \"world`
	expected := ParsedString{
		Raw:      input,
		Value:    `hello "world`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello "world"`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello world how" are you`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello "world`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello "world"`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello world how "are you`,
		Features: FullFeatures,
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
		Raw:      input,
		Value:    `hello "world how are you`,
		Features: FullFeatures,
	}

	actual := ParseRawString(input, FullFeatures)

	if !(cmp.Equal(expected, actual)) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
