// You can edit this code!
// Click here and start typing.
package sort


type Node struct {
	dist []int
	seen int
}

func BFS(G *map[int]*Node, v int) {
	var Q *Queue

	Q = NewQueue(10)

	(*G)[v].seen = 0
	Q.push(v)

	for !Q.isEmpty() {
		v = Q.pop()
		for _, i := range (*G)[v].dist {
			if (*G)[i].seen == -1 {
				(*G)[i].seen = (*G)[v].seen + 1
				Q.push(i)
			}
		}
	}
}