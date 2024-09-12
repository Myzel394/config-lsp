package common

func CursorToCharacterIndex(cursor uint32) uint32 {
	return max(0, cursor-1)
}
