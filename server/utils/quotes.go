package utils

import (
	"cmp"
	"slices"
)

type quoteRange [2]int

func (q quoteRange) IsCharInside(index int) bool {
	return index >= q[0] && index <= q[1]
}

type quoteRanges []quoteRange

func (q quoteRanges) IsIndexInsideQuotes(index int) bool {
	return q.GetQuoteForIndex(index) != nil
}

func (q quoteRanges) GetQuoteForIndex(index int) *quoteRange {
	index, found := slices.BinarySearchFunc(
		q,
		index,
		func(current quoteRange, target int) int {
			return cmp.Compare(target, current[0])
		},
	)

	if !found {
		return nil
	}

	return &q[index]
}

func GetQuoteRanges(s string) quoteRanges {
	quoteRanges := make(quoteRanges, 0, 2)
	inQuote := false
	var quoteStart int

	for index, c := range s {
		if c == '"' && (index == 0 || s[index-1] != '\\') {
			if inQuote {
				quoteRanges = append(quoteRanges, [2]int{quoteStart, index})
				inQuote = false
			} else {
				quoteStart = index
				inQuote = true
			}
		}
	}

	return quoteRanges
}
