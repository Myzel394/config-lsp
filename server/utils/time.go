package utils

import "time"

func BlockUntilNotNil(pointer any) {
	for pointer == nil {
		// This is a blocking call to wait until the pointer is not nil.
		// It can be used in scenarios where the pointer is expected to be set by another goroutine.
	}
}

func BlockUntilNotNilTimeout(pointer any, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)

	for pointer == nil {
		if time.Now().After(deadline) {
			return false
		}
		// This is a blocking call to wait until the pointer is not nil.
		// It can be used in scenarios where the pointer is expected to be set by another goroutine.
	}

	return true
}

// Waits till the provided pointer is not nil.
// Has a default timeout of 3 seconds
func BlockUntilIndexesNotNil(d any) bool {
	return BlockUntilNotNilTimeout(d, 3*time.Second)
}
