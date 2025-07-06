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

// Calculate a numeric value from a string and return it as a byte.
// Supports decimal values.
// Supported units:
// - `k` for kilobytes
// - `m` for megabytes
// - `g` for gigabytes
// - `t` for terabytes
// - `e` for exabytes
// - `p` for petabytes
// - `z` for zettabytes
// Supports bit / Byte suffixes:
// - `b` for bits
// - `B` for bytes
// For bits the value will be rounded down to the nearest byte.
// base is either 1000 or 1024.
// Returns an error if the value is invalid or the unit is not supported.
func CalculateNumericValueToByte(
	amount float64,
	unit rune,
	suffix string,
	base float64,
) (uint64, error) {
	if base != 1000 && base != 1024 {
		return 0, errors.New("base must be either 1000 or 1024")
	}

	if amount < 0 {
		return 0, errors.New("amount must be greater than or equal to 0")
	}

	var byteAmount uint64 = 0

	if suffix == "b" {
		// If the suffix is 'b', we treat it as bits.
		// 1 byte = 8 bits, so we divide the amount by 8.
		amount /= 8
	}

	switch unit {
	case 'k':
		byteAmount = uint64(amount * base)
	case 'm':
		byteAmount = uint64(amount * base*base)
	case 'g':
		byteAmount = uint64(amount * base*base*base)
	case 't':
		byteAmount = uint64(amount * base*base*base*base)
	case 'e':
		byteAmount = uint64(amount * base*base*base*base*base)
	case 'p':
		byteAmount = uint64(amount * base*base*base*base*base*base)
	case 'z':
		byteAmount = uint64(amount * base*base*base*base*base*base*base)
	default:
		byteAmount = uint64(amount)
	}

	if byteAmount < 0 {
		return 0, errors.New("calculated byte amount is negative")
	}

	return byteAmount, nil
}
