package utils

import (
	"errors"
	"fmt"
	"os"
)

func DoesPathExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func IsPathDirectory(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return info.IsDir()
}

func IsPathFile(path string) bool {
	_, err := os.Stat(path)

	print(fmt.Sprintf("Checking if path %s is a file: %v\n", path, err))

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
