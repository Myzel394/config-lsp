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
