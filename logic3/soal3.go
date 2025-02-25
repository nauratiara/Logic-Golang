package logic3

func Soal03(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			num := 2
			if i > 0 {
				num = matrix[i-1][i] + 2
			}
			for j := i; j < n; j++ {
				matrix[i][j] = num
				if i == 0 {
					num += 3
				} else {
					num += 2
				}
			}
		} else {
			num := 49
			if i > 1 {
				num = matrix[i-2][i] - 2
			}
			for j := n - 1; j >= i; j-- {
				matrix[i][j] = num
				num -= 3
			}
		}
	}
	return matrix
}
