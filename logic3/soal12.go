package logic3

func Soal12(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	nTengah := (n - 1) / 2
	num := 1

	for row := 0; row <= nTengah; row++ {
		for col := 0; col < n; col++ {
			if (row == 0 || row == nTengah) && (col == row || col == n-1-row) {
				matrix[row][col] = num
				matrix[n-1-row][col] = num
			} else if row > 0 && row < nTengah {
				if col == row {
					matrix[row][col] = num + 4*row
					matrix[n-1-row][col] = num + 4*row
				} else if col == row+1 {
					matrix[row][col] = num + 4*row - 2
					matrix[n-1-row][col] = num + 4*row - 2
				} else if col == n-1-row {
					matrix[row][col] = num + 2*row
					matrix[n-1-row][col] = num + 2*row
				} else if col > nTengah && (col-row) <= nTengah && (col+row) >= nTengah+1 && (col+row) <= n+(nTengah-row)-1 {
					matrix[row][col] = num + 2*(n-1-col)
					matrix[n-1-row][col] = num + 2*(n-1-col)
				}
			}
		}
	}
	return matrix
}
