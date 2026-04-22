package strings_problems

import "strings"

// RemoveOddIndexChars demonstrates index-based filtering with strings.Builder for O(n) construction.
func RemoveOddIndexChars(s string) string {
	var b strings.Builder
	for i, ch := range s {
		if i%2 == 0 {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
