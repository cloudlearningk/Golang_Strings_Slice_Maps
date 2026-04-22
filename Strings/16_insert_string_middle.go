package strings_problems

// InsertStringMiddle shows mid-point slicing — splitting outer at len/2 assumes even length per spec.
func InsertStringMiddle(outer, inner string) string {
	mid := len(outer) / 2
	return outer[:mid] + inner + outer[mid:]
}
