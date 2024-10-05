package utils

import "slices"

type quoteRange [2]int

func (q quoteRange) IsCharInside(index int) bool {
	return index >= q[0] && index <= q[1]
}

type quoteRanges []quoteRange

func (q quoteRanges) IsCharInside(index int) bool {
	_, found := slices.BinarySearchFunc(
		q,
		index,
		func(current quoteRange, target int) int {
			if target < current[0] {
				return -1
			}

			if target > current[1] {
				return 1
			}

			return 0
		},
	)

	return found
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
