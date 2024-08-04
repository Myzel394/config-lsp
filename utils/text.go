package utils

func CountCharacterOccurrences(line string, character rune) int {
	count := 0

	for _, c := range line {
		if c == character {
			count++
		}
	}

	return count
}
