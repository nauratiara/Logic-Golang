package logic2

func Soal13(n int) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row <= col && row+col <= n-1 {
				matrix[row][col] = num
				matrix[n-1-row][col] = num
			}
			num += 2
		}
	}
	return matrix
}
