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

func (q quoteRanges) GetInvertedRanges(textLength int) [][2]int {
	if textLength == 0 {
		return nil
	}

	if len(q) == 0 {
		return [][2]int{
			{
				0,
				textLength - 1,
			},
		}
	}

	inverted := make([][2]int, 0, len(q))

	firstRange := q[0]

	if firstRange[0] != 0 {
		inverted = append(inverted, [2]int{0, firstRange[0]})
	}

	if len(q) == 1 {
		return inverted
	}

	for index, currentRange := range q[:len(q)-1] {
		nextRange := q[index+1]

		inverted = append(inverted, [2]int{currentRange[1] + 1, nextRange[0]})
	}

	lastRange := q[len(q)-1]

	if lastRange[1] != (textLength - 1) {
		inverted = append(inverted, [2]int{lastRange[1], textLength - 1})
	}

	return inverted
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
