package strings_problems

// FirstLastTwo demonstrates edge-case handling: short strings are a common interview gotcha.
func FirstLastTwo(s string) string {
	n := len(s)
	if n < 2 {
		return ""
	}
	if n == 2 {
		return s + s
	}
	return s[:2] + s[n-2:]
}
