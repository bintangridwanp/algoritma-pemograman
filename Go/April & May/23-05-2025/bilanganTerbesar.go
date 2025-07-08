package main

import "fmt"

func main() {
	angka := []int{3, 1, 4, 1}

	// Urutkan angka dari terbesar ke terkecil (sederhana)
	for i := 0; i < len(angka); i++ {
		for j := i + 1; j < len(angka); j++ {
			if angka[i] < angka[j] {
				// Tukar posisi
				temp := angka[i]
				angka[i] = angka[j]
				angka[j] = temp
			}
		}
	}

	// Cetak angka sebagai satu bilangan
	fmt.Print("Bilangan terbesar: ")
	for i := 0; i < len(angka); i++ {
		fmt.Print(angka[i])
	}
	fmt.Println()
}
