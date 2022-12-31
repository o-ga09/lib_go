// You can edit this code!
// Click here and start typing.
package sort

import "lib/std"

func quickSort(a *[]int, p int, r int) {
	if p < r {
		q := parttion(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}
}

func parttion(a *[]int, p int, r int) int {
	x := (*a)[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if (*a)[j] <= x {
			i++
			std.Swap(a, i, j)
		}
	}
	std.Swap(a, i+1, r)
	return i + 1
}