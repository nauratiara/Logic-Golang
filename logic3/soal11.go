package logic3

func Soal11(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	start := 1
	mid := (n - 1) / 2
	for i := 0; i < mid; i++ {
		matrix[i][i] = start
		matrix[n-1-i][i] = start
		start += 2
	}

	start = 1
	for i := 0; i < mid; i++ {
		if i%2 == 0 {
			for j := n - 1; j >= mid; j-- {
				if i+j >= 8 {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
					start += 2
				}
			}
		} else {
			for j := mid; j < n; j++ {
				if i+j >= 8 {
					matrix[i][j] = start
					matrix[n-1-i][j] = start
					start += 2
				}
			}
		}
	}
	return matrix
}
