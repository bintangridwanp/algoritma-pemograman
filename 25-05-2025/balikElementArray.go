package main

import "fmt"

func main() {
	angka := []int{1, 2, 3, 4, 5, 6}
	hasil := []int{}

	// Membalikkan elemen array
	for i := len(angka) - 1; i >= 0; i-- {
		hasil = append(hasil, angka[i])
	}

	// Menampilkan hasil
	for i := 0; i < len(hasil); i++ {
		fmt.Print(hasil[i], " ")
	}
}
