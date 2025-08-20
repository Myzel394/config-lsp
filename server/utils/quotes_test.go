package utils

import "testing"

func TestQuotesSimpleExample(
	t *testing.T,
) {
	quoteRanges := GetQuoteRanges(`"hello" "world"`)

	if !(len(quoteRanges) == 2 && quoteRanges[0][0] == 0 && quoteRanges[0][1] == 6 && quoteRanges[1][0] == 8 && quoteRanges[1][1] == 14) {
		t.Fatalf("Unexpected quote ranges: %v", quoteRanges)
	}
}

func TestQuotesEscapedQuotes(
	t *testing.T,
) {
	quoteRanges := GetQuoteRanges(`"hello \"world\""`)

	if !(len(quoteRanges) == 1 && quoteRanges[0][0] == 0 && quoteRanges[0][1] == 16) {
		t.Fatalf("Unexpected quote ranges: %v", quoteRanges)
	}
}

func TestInvertedQuotesSimple(
	t *testing.T,
) {
	text := `"hello" "world"`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 1 && inverted[0][0] == 7 && inverted[0][1] == 8) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestInvertedQuotesImmediatelyFollowing(
	t *testing.T,
) {
	text := `"hello""world"`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 1 && inverted[0][0] == 7 && inverted[0][1] == 7) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestInvertedQuotesEscapedQuotes(
	t *testing.T,
) {
	text := `hello \"world\"`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 1 && inverted[0][0] == 0 && inverted[0][1] == 14) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestInvertedQuotesMultiple(
	t *testing.T,
) {
	text := `"hello" "world" "hello"`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 2 && inverted[0][0] == 7 && inverted[0][1] == 8 && inverted[1][0] == 15 && inverted[1][1] == 16) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestInvertedQuotesFullyQuoted(
	t *testing.T,
) {
	text := `"hello world"`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 0) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestInvertedQuotesFirstThenRemaining(
	t *testing.T,
) {
	text := `"hello world" i am here`
	quoteRanges := GetQuoteRanges(text)
	inverted := quoteRanges.GetInvertedRanges(len(text))

	if !(len(inverted) == 1 && inverted[0][0] == 13 && inverted[0][1] == 23) {
		t.Fatalf("Unexpected inverted quote ranges: %v", inverted)
	}
}

func TestNormalizeValidQuotes(
	t *testing.T,
) {
	// Test with valid quotes
	input := `hello \"world"`
	expected := `hello "world"`
	result := NormalizeEscapedQuotes(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
