package logic3

func Soal10(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	nTengah := (n - 1) / 2
	num := 9

	for row := 0; row <= nTengah; row++ {
		for col := 0; col < n; col++ {
			if col+row >= nTengah && col-row <= nTengah {
				matrix[row][col] = num
				matrix[n-1-row][col] = num
				num -= 2
			}
		}
		num = 9
	}
	return matrix
}
