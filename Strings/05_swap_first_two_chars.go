package strings_problems

import "fmt"

// SwapFirstTwo illustrates string slicing for in-place-style swaps without mutation.
func SwapFirstTwo(a, b string) string {
	if len(a) < 2 || len(b) < 2 {
		return fmt.Sprintf("%s %s", a, b)
	}
	newA := b[:2] + a[2:]
	newB := a[:2] + b[2:]
	return newA + " " + newB
}
