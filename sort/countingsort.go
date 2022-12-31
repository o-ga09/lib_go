// You can edit this code!
// Click here and start typing.
package sort

import "lib/std"

func countingSort(a *[]int, k int) *[]int {
	var b []int
	var c []int

	b = make([]int, std.N+1)
	c = make([]int, std.N+1)
	for i := 0; i <= std.N; i++ {
		c[i] = 0
	}

	for j := 0; j < len((*a)); j++ {
		c[(*a)[j]] = c[(*a)[j]] + 1
	}

	for i := 0; i < std.N; i++ {
		c[i+1] = c[i+1] + c[i]
	}

	for j := len((*a)) - 1; j >= 0; j-- {
		b[c[(*a)[j]]] = (*a)[j]
		c[(*a)[j]] = c[(*a)[j]] - 1
	}
	return &b
}
