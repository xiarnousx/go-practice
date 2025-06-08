// Package testexample shows docs and test examples
package testexample

func Abc(s string) string {
	return s
}

// Sum sum up int arguments
func Sum(n ...int) int {
	var s int
	for _, v := range n {
		s += v
	}

	return s
}
