package strings_problems

import "strings"

// ReplaceNotPoor shows substring index comparison — key technique for positional string logic.
func ReplaceNotPoor(s string) string {
	notIdx := strings.Index(s, "not")
	poorIdx := strings.Index(s, "poor")
	if notIdx == -1 || poorIdx == -1 || notIdx > poorIdx {
		return s
	}
	return s[:notIdx] + "good" + s[poorIdx+4:]
}
