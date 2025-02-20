package logic2

func Soal01(n int) (matrix [][]int) {
	matrix = make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = 2*j + 1
		}
	}

	return matrix
}
