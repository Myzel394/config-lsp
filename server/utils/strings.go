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

func FindFirstNonMatch(s string, substr map[rune]struct{}, startIndex int) int {
	for index := startIndex; index < len(s); index++ {
		if _, found := substr[rune(s[index])]; !found {
			return index
		}
	}

	return -1
}

func FindLastNonMatch(s string, substr map[rune]struct{}, startIndex int) int {
	for index := startIndex; index >= 0; index-- {
		if _, found := substr[rune(s[index])]; !found {
			return index
		}
	}

	return -1
}

var lineContinuationPattern = regexp.MustCompile(`\\\s*$`)

// SplitSmartlyIntoLines Split a string into lines while respecting "\" as a line continuation character
// This function is useful for parsing configuration files
// You will need to handle whitespace trimming yourself
// For example, the following input:
// ```
// key1 = value1
//
//	key2 = value2 \
//	    value3
//
// key3 = value4
// ```
// Will be split into:
// ```go
//
//	[][]string{
//		[]string{"key1 = value1"},
//		[]string{"key2 = value2 ", "    value3"},
//		[]string{"key3 = value4"},
//	}
func SplitIntoVirtualLines(input string) [][]string {
	lines := make([][]string, 0, len(input))
	currentLine := make([]string, 0, 1)

	for _, line := range SplitIntoLines(input) {
		if lineContinuationPattern.MatchString(line) {
			currentLine = append(currentLine, line[:len(line)-1])
			continue
		}

		currentLine = append(currentLine, line)
		lines = append(lines, currentLine)
		currentLine = make([]string, 0, 1)
	}

	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}
