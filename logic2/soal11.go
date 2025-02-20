package logic2

func Soal11(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		value := 1
		for i := 0; i <= j; i++ {
			matrix[i][j] = value
			value += 2
		}
	}

	return matrix
}
