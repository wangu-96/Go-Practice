package strutil

func Reverse(s string) string {
	// Convert the string to a slice of runes to handle Unicode characters.
	runes := []rune(s)

	// Iterate with two pointers, swapping elements from the start and end.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string.
	return string(runes)
}
