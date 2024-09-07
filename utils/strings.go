package utils

import (
	"regexp"
)

var trimIndexPattern = regexp.MustCompile(`^\s*(.+?)\s*$`)

func GetTrimIndex(s string) []int {
	indexes := trimIndexPattern.FindStringSubmatchIndex(s)

	if indexes == nil {
		return nil
	}

	return indexes[2:4]
}

func SplitIntoLines(s string) []string {
	return regexp.MustCompile("\r?\n").Split(s, -1)
}

func FindPreviousCharacter(line string, character string, startIndex int) (int, bool) {
	for index := startIndex; index >= 0; index-- {
		if string(line[index]) == character {
			return index, true
		}
	}

	return 0, false
}

func FindNextCharacter(line string, character string, startIndex int) (int, bool) {
	for index := startIndex; index < len(line); index++ {
		if string(line[index]) == character {
			return index, true
		}
	}

	return 0, false
}

func CountCharacterOccurrences(line string, character rune) int {
	count := 0

	for _, c := range line {
		if c == character {
			count++
		}
	}

	return count
}
