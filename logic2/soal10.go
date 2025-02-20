package logic2

func Soal10(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		value := 1
		for j := 0; j <= i; j++ {
			matrix[i][j] = value
			value += 2
		}
	}

	return matrix
}
