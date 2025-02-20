package logic2

func Soal05(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][n-1-j] = start
				start += 2
			}
		}
	}

	return result
}
