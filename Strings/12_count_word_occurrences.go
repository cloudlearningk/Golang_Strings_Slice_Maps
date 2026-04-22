package strings_problems

import "strings"

// WordOccurrences uses strings.Fields to split on any whitespace — more robust than Split(" ").
func WordOccurrences(s string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	return counts
}
