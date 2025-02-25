package logic3

func Soal07(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	nTengah := (n - 1) / 2
	num := 1
	for row := 0; row <= n/2; row++ {
		for col := 0; col < n; col++ {
			if col+row >= nTengah && col-row <= nTengah {
				if row%2 == 1 {
					matrix[row][col] = num
					matrix[n-1-row][col] = num
				} else {
					matrix[row][n-1-col] = num
					matrix[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return matrix
}
