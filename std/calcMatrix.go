// You can edit this code!
// Click here and start typing.
package std

// func main() {

// }

func CalcMatrix(x *[][]int, y *[][]int) *[][]int {

	var res [][]int
	
	N := len((*x))
	M := len((*y))
	
	res = make([][]int,N,10)
	for i := 0; i < N; i++ {
		res[i] = make([]int,M,10)
	}
		
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			res[i][j] += (*x)[i][j] * (*y)[j][i]
		}
	}

	return &res
}