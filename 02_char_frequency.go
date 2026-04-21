package strings_problems

// CharFrequency uses a map to count rune frequencies — preferred over byte counting for unicode correctness.
func CharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}
	return freq
}
