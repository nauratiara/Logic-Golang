package logic1

import (
	"fmt"
)

func Soal1(n int) []int {
	fmt.Printf("Soal No. 1\n")
	fmt.Printf("n = %d\n\n", n)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 2*i + 1
	}

	return result
}

func Soal2(n int) []int {
	fmt.Printf("Soal No. 2\n")
	fmt.Printf("n = %d\n\n", n)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 2 * (i + 1)
	}
	return result
}

func Soal3(n int) []int {
	fmt.Printf("Soal No. 3\n")
	fmt.Printf("n = %d\n\n", n)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 3 * (i + 1)
	}
	return result
}

func Soal4(n int) []int {
	fmt.Printf("Soal No. 4\n")
	fmt.Printf("n = %d\n\n", n)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 2*(n-i) - 1
	}
	return result
}
