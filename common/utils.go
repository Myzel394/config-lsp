package common

import (
	"os"
	"strings"
)

func GetLine(path string, line int) (string, error) {
	path = path[len("file://"):]
	readBytes, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	// Split file into lines
	lines := strings.Split(string(readBytes), "\n")

	return lines[line], nil
}

func Map[T any, O any](values []T, f func(T) O) []O {
	result := make([]O, len(values))

	for index, value := range values {
		result[index] = f(value)
	}

	return result
}
