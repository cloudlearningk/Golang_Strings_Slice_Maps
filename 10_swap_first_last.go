package strings_problems

// SwapFirstLast shows byte-level manipulation — valid here since we only touch positions 0 and n-1.
func SwapFirstLast(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	b := []byte(s)
	b[0], b[n-1] = b[n-1], b[0]
	return string(b)
}
