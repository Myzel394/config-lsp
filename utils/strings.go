package utils

import "strings"

func IndexOffset(s string, search string, start int) int {
	return strings.Index(s[start:], search) + start
}
