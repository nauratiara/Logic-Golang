package logic2

func Soal07(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	value := 1
	for i := 0; i < n; i++ {
		matrix[i][i] = value
		value += 2
	}

	return matrix
}
