package logic3

func Soal12(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	start := 1
	mid := (n - 1) / 2
	for i := 0; i < mid; i++ {
		matrix[i][n-1-i] = start
		matrix[n-1-i][n-1-i] = start
		start += 2
	}

	start = 1
	for i := 0; i < mid; i++ {
		if i%2 == 0 {
			for j := 0; j <= mid; j++ {
				if i >= j {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
					start += 2
				}
			}
		} else {
			for j := mid; j >= 0; j-- {
				if i >= j {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
					start += 2
				}
			}
		}
	}
	return matrix
}
