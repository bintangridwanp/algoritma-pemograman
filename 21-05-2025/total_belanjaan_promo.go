package main

import "fmt"

func main() {
	hargaBarang := []int{30000, 60000, 80000, 50000}
	totalBayar := 0
	tempatKosong := 0

	// Loop untuk menghitung total belanjaan dengan promo
	for i := 0; i < len(hargaBarang); i++ {
		// Jika harga barang lebih dari 50.000, diskon 10%
		if hargaBarang[i] > 50000 {
			tempatKosong = hargaBarang[i] - 10000
			totalBayar = totalBayar + tempatKosong
			tempatKosong = 0
		} else {
			// Jika harga barang kurang 50.000, tidak ada diskon.
			totalBayar = totalBayar + hargaBarang[i]
		}
	}
	fmt.Println("Total Bayar", totalBayar)
}
