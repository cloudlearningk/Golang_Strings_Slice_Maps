package strings_problems

import "strings"

// ReplaceFirstChar shows targeted replacement: only substitute occurrences after index 0.
func ReplaceFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	first := s[0]
	rest := s[1:]
	return string(first) + strings.ReplaceAll(rest, string(first), "$")
}
