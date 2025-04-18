package utils

import (
	"regexp"
	"strings"
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

var emptyRegex = regexp.MustCompile(`^\s*$`)

func IsEmpty(s string) bool {
	return emptyRegex.MatchString(s)
}

func AllIndexes(s string, sub string) []int {
	indexes := make([]int, 0)
	current := s

	for {
		index := strings.Index(current, sub)

		if index == -1 {
			break
		}

		indexes = append(indexes, index)
		current = current[index+1:]
	}

	return indexes
}
