package strings_problems

import "strings"

// RepeatLastTwo uses strings.Repeat to avoid manual loops — idiomatic for fixed-count repetition.
func RepeatLastTwo(s string) string {
	if len(s) < 2 {
		return s
	}
	last2 := s[len(s)-2:]
	return strings.Repeat(last2, 4)
}
