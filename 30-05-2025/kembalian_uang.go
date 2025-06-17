package main

import "fmt"

func main() {
	var bayar, hargaBarang int
	fmt.Print("Enter your Harga Barang: ")
	fmt.Scanln(&hargaBarang)
	fmt.Print("Masukkan Uang Bayar: ")
	fmt.Scanln(&bayar)

	hasil := bayar - hargaBarang

	if hasil > 0 {
		fmt.Println("Kembalian Uang: ", hasil)
	} else {
		fmt.Println("Uang tidak cukup untuk membayar barang")
	}
}
