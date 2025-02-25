package logic3

func Soal08(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	nTengah := (n - 1) / 2
	num := 1
	for col := 0; col <= n/2; col++ {
		for row := 0; row < n; row++ {
			if col+row >= nTengah && row-col <= nTengah {
				if col%2 == 1 {
					matrix[row][col] = num
					matrix[row][n-1-col] = num
				} else {
					matrix[n-1-row][col] = num
					matrix[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return matrix
}
