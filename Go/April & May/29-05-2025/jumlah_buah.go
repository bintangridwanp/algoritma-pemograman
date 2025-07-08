package main

import "fmt"

func main() {

	// Buat slice untuk menyimpan jumlah buah
	jumlahBuah := []int{10, 20, 30, 40, 50}

	// Hitung total jumlah buah
	total := 0
	for i := 0; i < len(jumlahBuah); i++ {
		total += jumlahBuah[i]
	}
	// Cetak total jumlah buah
	fmt.Println("Total jumlah buah:", total)

}
