package utils

import (
	"regexp"
	"strings"
)

func IndexOffset(s string, search string, start int) int {
	return strings.Index(s[start:], search) + start
}

var trimIndexPattern = regexp.MustCompile(`^\s+(.+?)\s+`)

func GetTrimIndex(s string) []int {
	indexes := trimIndexPattern.FindStringSubmatchIndex(s)

	if indexes == nil {
		return nil
	}

	return indexes[2:4]
}
