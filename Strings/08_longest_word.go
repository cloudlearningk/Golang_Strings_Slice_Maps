package strings_problems

// LongestWord uses multiple return values — idiomatic Go for returning both result and metadata.
func LongestWord(words []string) (string, int) {
	if len(words) == 0 {
		return "", 0
	}
	longest := words[0]
	for _, w := range words[1:] {
		if len(w) > len(longest) {
			longest = w
		}
	}
	return longest, len(longest)
}
