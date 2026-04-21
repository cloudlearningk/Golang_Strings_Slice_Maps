package strings_problems

import "unicode/utf8"

// StringLength returns both byte length and character (rune) length of a string.
// In Go, len(s) counts bytes, not characters. For ASCII strings they match,
// but for Unicode strings (e.g., emojis, Chinese chars) they differ.
func StringLength(s string) (byteLen int, charLen int) {
	byteLen = len(s)                    // O(1) — stored in string header
	charLen = utf8.RuneCountInString(s) // O(n) — iterates bytes
	return
}
