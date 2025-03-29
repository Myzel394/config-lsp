package common

import (
	"github.com/hbollon/go-edlib"
)

// Find items that are similar to the given input.
// This is used to find typos & suggest the correct item.
// Once an item is found that has a Damerau-Levenshtein distance of 1, it is immediately returned.
// If not, then the next 2 items of similarity 2, or 3 items of similarity 3 are returned.
// If no items with similarity <= 3 are found, then an empty slice is returned.
func FindSimilarItems[T ~string](
	input T,
	items []T,
) []T {
	itemsPerSimilarity := map[uint8][]T{
		2: make([]T, 0, 2),
		3: make([]T, 0, 3),
	}

	for _, item := range items {
		similarity := edlib.DamerauLevenshteinDistance(string(item), string(input))

		switch similarity {
		case 1:
			return []T{item}
		case 2:
			itemsPerSimilarity[2] = append(itemsPerSimilarity[2], item)

			if len(itemsPerSimilarity[2]) >= 2 {
				return itemsPerSimilarity[2]
			}
		case 3:
			itemsPerSimilarity[3] = append(itemsPerSimilarity[3], item)

			if len(itemsPerSimilarity[3]) >= 3 {
				return itemsPerSimilarity[3]
			}
		}
	}

	return append(itemsPerSimilarity[2], itemsPerSimilarity[3]...)
}
