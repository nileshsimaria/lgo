// Package mystringutil contains utility functions for working with strings.
package mystringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1
	for i < len(r)/2 {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
