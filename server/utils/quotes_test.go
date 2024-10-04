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
