package logic2

func Soal02(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 1; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = 2 * i
		}
	}
	return matrix
}
