package logic2

func Soal04(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	start := 1

	for i := 1; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = start
			start += 3
		}
	}
	return matrix
}
