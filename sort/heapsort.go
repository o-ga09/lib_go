// You can edit this code!
// Click here and start typing.
package sort

import "lib/std"

func heapSort(a *[]int) {
	for i := len((*a)) / 2; i >= 0; i-- {
		downheap(a, i, len((*a))-1)
	}

	for i := len((*a)) - 1; i >= 1; i-- {
		std.Swap(a, 0, i)
		downheap(a, 0, i-1)
	}

}

func downheap(a *[]int, i int, n int) {
	var j int

	for {
		j = 2*i + 1
		if j > n {
			break
		}
		if j < n && (*a)[j+1] > (*a)[j] {
			j++
		}
		if (*a)[i] >= (*a)[j] {
			break
		}
		std.Swap(a, i, j)
		i = j
	}
}