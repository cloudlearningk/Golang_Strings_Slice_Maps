package strings_problems

// ReverseIfMultipleOf4 combines modular arithmetic check with rune-level reversal for unicode safety.
func ReverseIfMultipleOf4(s string) string {
	if len(s)%4 != 0 {
		return s
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
