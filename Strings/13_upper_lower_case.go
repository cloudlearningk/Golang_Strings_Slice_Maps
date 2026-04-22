package strings_problems

import "strings"

// UpperAndLower uses named returns to make the dual-output intent explicit at the call site.
func UpperAndLower(s string) (upper, lower string) {
	upper = strings.ToUpper(s)
	lower = strings.ToLower(s)
	return
}
