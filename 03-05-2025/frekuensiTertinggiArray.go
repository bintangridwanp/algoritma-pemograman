package main

import "fmt"

func main() {
	data := [...]int{1, 2, 2, 3, 3, 3, 4, 2}

	var frekuensi [101]int // Asumsikan nilai dalam array <= 100

	// Hitung frekuensi setiap nilai
	for i := 0; i < len(data); i++ {
		frekuensi[data[i]]++
	}

	// Cari nilai frekuensi maksimum
	maks := 0
	for i := 0; i < len(frekuensi); i++ {
		if frekuensi[i] > maks {
			maks = frekuensi[i]
		}
	}

	// Cetak semua nilai dengan frekuensi tertinggi
	fmt.Println("Nilai paling sering muncul:")
	for i := 0; i < len(frekuensi); i++ {
		if frekuensi[i] == maks {
			fmt.Printf("%d (muncul %d kali)\n", i, maks)
		}
	}
}
