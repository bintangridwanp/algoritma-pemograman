package main

import "fmt"

func main() {
	nilaiSiswa := []int{75, 82, 85, 90, 70, 60, 95, 90, 91, 78}
	a := 0
	maksBerturut := 0

	for i := 0; i < len(nilaiSiswa); i++ {
		if nilaiSiswa[i] > 80 {
			a++
			if a > maksBerturut {
				maksBerturut = a
			}
		} else {
			a = 0
		}
	}

	fmt.Printf("Jumlah nilai berturut-turut tertinggi di atas 80: %d\n", maksBerturut)
}
