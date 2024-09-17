package common

func CursorToCharacterIndex(cursor uint32) uint32 {
	return max(1, cursor) - 1
}
