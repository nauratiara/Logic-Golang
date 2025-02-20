package logic1

func Soal5(n int) []int {
	result := make([]int, n)
	value := 20

	for i := 0; i < n; i++ {
		result[i] = value
		value -= 2
	}

	return result
}
