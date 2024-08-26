package utils

func Dedent(s string) string {
	return s[len("\n"):]
}

func KeyExists[T comparable, V any](keys map[T]V, key T) bool {
	_, found := keys[key]

	return found
}
