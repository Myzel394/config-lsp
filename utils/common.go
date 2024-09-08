package utils

import (
	"os"
)

func DoesPathExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func IsPathDirectory(path string) bool {
	info, err := os.Stat(path)

	return err == nil && info.IsDir()
}

func IsPathFile(path string) bool {
	info, err := os.Stat(path)

	return err == nil && !info.IsDir()
}
