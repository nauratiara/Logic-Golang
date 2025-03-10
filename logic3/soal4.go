package logic3

func Soal04(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	num := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := n - 1; j >= n-1-i; j-- {
				matrix[i][j] = num
				num += 2
			}
		} else {
			for j := n - 1 - i; j < n; j++ {
				matrix[i][j] = num
				num += 2
			}
		}
	}

	return matrix
}
