package logic2

func Soal12(n int) (matrix [][]int) {
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

	for i := 0; i <= n/2; i++ {
		value := 17
		for j := n - 1; j >= n-1-i; j-- {
			matrix[i][j] = value
			value -= 2
		}
	}
	for i := n / 2; i < n; i++ {
		value := 17
		for j := n - 1; j >= n-1-(n-1-i); j-- {
			matrix[i][j] = value
			value -= 2
		}
	}

	return matrix
}
