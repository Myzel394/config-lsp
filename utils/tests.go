package utils

import "strings"

func Dedent(s string) string {
	return strings.TrimLeft(s, "\n")
}

func KeyExists[T comparable, V any](keys map[T]V, key T) bool {
	_, found := keys[key]

	return found
}
