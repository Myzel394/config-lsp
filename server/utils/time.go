package utils

import "time"

func BlockUntilNotNil(pointer any) {
	for pointer == nil {
		// This is a blocking call to wait until the pointer is not nil.
		// It can be used in scenarios where the pointer is expected to be set by another goroutine.
		time.Sleep(10 * time.Millisecond)
	}
}

func BlockUntilNotNilTimeout(pointer any, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	sleepInterval := 10 * time.Millisecond

	for pointer == nil {
		if time.Now().After(deadline) {
			return false
		}
		// Sleep to avoid busy waiting and reduce CPU usage
		time.Sleep(sleepInterval)
	}

	return true
}

// Waits till the provided pointer is not nil.
// Has a default timeout of 3 seconds
func BlockUntilIndexesNotNil(d any) bool {
	return BlockUntilNotNilTimeout(d, 3*time.Second)
}
