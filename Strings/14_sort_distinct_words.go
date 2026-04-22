package strings_problems

import (
	"sort"
	"strings"
)

// SortDistinctWords uses a map for deduplication then sort.Strings — the standard Go dedup pattern.
func SortDistinctWords(s string) string {
	seen := make(map[string]bool)
	var distinct []string
	for _, w := range strings.Split(s, ",") {
		word := strings.TrimSpace(w)
		if word != "" && !seen[word] {
			seen[word] = true
			distinct = append(distinct, word)
		}
	}
	sort.Strings(distinct)
	return strings.Join(distinct, ",")
}
