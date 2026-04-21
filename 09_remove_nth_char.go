package strings_problems

// RemoveNthChar uses slice concatenation — in Go strings are immutable so we build a new one.
func RemoveNthChar(s string, n int) string {
	if n < 0 || n >= len(s) {
		return s
	}
	return s[:n] + s[n+1:]
}
