package utils

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSplitIntoVirtualLinesSimpleExample(
	t *testing.T,
) {
	input := Dedent(`
Hello
World\
how are you
`)
	expected := [][]string{
		{"Hello"},
		{"World", "how are you"},
	}

	actual := SplitIntoVirtualLines(input)

	if cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesEmptyString(t *testing.T) {
	input := ""
	expected := [][]string{
		{""},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesSingleLine(t *testing.T) {
	input := Dedent(`
    Hello`)
	expected := [][]string{
		{"    Hello"},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesMultipleLinesWithoutContinuation(t *testing.T) {
	input := Dedent(`
    Hello
    World
    How are you`)
	expected := [][]string{
		{"    Hello"},
		{"    World"},
		{"    How are you"},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesMultipleLinesWithContinuation(t *testing.T) {
	input := Dedent(`
    Hello \
World \
How are you`)
	expected := [][]string{
		{"    Hello ", "World ", "How are you"},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesMixedContinuation(t *testing.T) {
	input := Dedent(`
Hello
World\
    How are you`)
	expected := [][]string{
		{"Hello"},
		{"World", "    How are you"},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

func TestSplitIntoVirtualLinesTrailingContinuation(t *testing.T) {
	input := Dedent(`
Hello\
    `)
	expected := [][]string{
		{"Hello", "    "},
	}

	actual := SplitIntoVirtualLines(input)

	if !cmp.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
