// You can edit this code!
// Click here and start typing.
package sort

import "lib/std"

func bubbleSort(a *[]int) {
	for i := 0; i < len(*a)-1; i++ {
		for j := 1; j < len(*a)-i; j++ {
			if (*a)[j] < (*a)[j-1] {
				std.Swap(a, j, j-1)
			}
		}
	}
}
