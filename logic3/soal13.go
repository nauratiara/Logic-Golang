package logic3

func Soal13(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	mid := (n - 1) / 2

	for i := 0; i <= mid; i++ {
		start := 1
		for j := 0; j <= mid; j++ {
			if i+j >= mid {
				if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
					matrix[i][n-1-j] = start
					matrix[n-1-i][n-1-j] = start
				}
			}
			start += 2
		}
	}
	return matrix
}
