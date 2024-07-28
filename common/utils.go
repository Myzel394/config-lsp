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
