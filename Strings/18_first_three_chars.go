package strings_problems

// FirstThree shows guard clause pattern — check length before slicing to avoid panic.
func FirstThree(s string) string {
	if len(s) <= 3 {
		return s
	}
	return s[:3]
}
