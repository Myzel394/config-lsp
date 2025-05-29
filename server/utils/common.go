package utils

import (
	"errors"
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

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
