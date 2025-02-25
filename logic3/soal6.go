package logic3

func Soal06(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for row := 0; row < n/2+1; row++ { // Hanya iterasi setengah baris
		num := 1
		if row%2 == 1 {
			num = (n-row-1)*2 + 1
		}
		for col := 0; col < n; col++ {
			if row <= col && row+col <= n-1 {
				matrix[row][col] = num
				matrix[n-1-row][col] = num
				if row%2 == 0 {
					num += 2
				} else {
					num -= 2
				}
			} else if row <= col {
				if row%2 == 0 {
					num += 2
				} else {
					num -= 2
				}
			}
		}
	}
	return matrix
}
