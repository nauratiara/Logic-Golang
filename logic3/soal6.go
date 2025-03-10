package logic3

func Soal06(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i <= j {
				if i%2 == 0 {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
				} else {
					matrix[i][n-1-j] = start
					matrix[n-1-i][n-1-j] = start
				}
				start += 2
			}
		}
	}
	return matrix
}
