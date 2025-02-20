package logic1

func Soal6(n int) []int {
	result := make([]int, n)
	value := 30

	for i := 0; i < n; i++ {
		result[i] = value
		value -= 3
	}

	return result
}
