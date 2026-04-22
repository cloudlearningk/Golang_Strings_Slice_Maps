package strings_problems

import "strings"

// SubstringBeforeChar uses strings.LastIndexByte — more efficient than scanning manually for last occurrence.
func SubstringBeforeChar(s string, ch byte) string {
	idx := strings.LastIndexByte(s, ch)
	if idx == -1 {
		return s
	}
	return s[:idx]
}
