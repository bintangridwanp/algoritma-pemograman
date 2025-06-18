package main

import "fmt"

func main() {
	angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	var total int

	for i := 0; i < len(angka); i++ {
		total += angka[i]
	}

	rataRata := total / len(angka)
	fmt.Println("Rata-rata:", rataRata)
}
