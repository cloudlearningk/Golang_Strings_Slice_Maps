package strings_problems

import "strings"

// AddIngOrLy demonstrates suffix detection with strings.HasSuffix — cleaner than manual slicing.
func AddIngOrLy(s string) string {
	if len(s) < 3 {
		return s
	}
	if strings.HasSuffix(s, "ing") {
		return s + "ly"
	}
	return s + "ing"
}
