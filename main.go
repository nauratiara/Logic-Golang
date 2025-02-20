package main

import "fmt"

func soal1(n int) {
	fmt.Printf("Soal No. 1\n")
	fmt.Printf("n = %d\n\n", n)
	fmt.Printf(" ")

	for i := 1; i <= n; i++ {
		fmt.Print(2*i-1, " ")
	}

	fmt.Println()
}

func soal2(n int) {
	fmt.Printf("Soal No. 2\n")
	fmt.Printf("n = %d\n\n", n)
	fmt.Printf(" ")

	for i := 1; i <= n; i++ {
		fmt.Print(2*i, " ")
	}

	fmt.Println()
}

func soal3(n int) {
	fmt.Printf("Soal No. 3\n")
	fmt.Printf("n = %d\n\n", n)
	fmt.Printf(" ")

	for i := 1; i <= n; i++ {
		fmt.Print(3*i, " ")
	}

	fmt.Println()
}

func soal4(n int) {
	fmt.Printf("Soal No. 4\n")
	fmt.Printf("n = %d\n\n", n)
	fmt.Printf(" ")

	for i := n; i >= 1; i-- {
		fmt.Print(2*i-1, " ")
	}

	fmt.Println()
}

func main() {
	n := 10

	soal1(n)
	fmt.Println()

	soal2(n)
	fmt.Println()

	soal3(n)
	fmt.Println()

	soal4(n)
}
