package logic1

import "fmt"

func Soal2(n int) []int {
	fmt.Printf("Soal No. 2\n")
	fmt.Printf("n = %d\n\n", n)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 2 * (i + 1)
	}
	return result
}
