package main

import "fmt"

func main() {

	barang := [...]int{1, 0, 1, 1, 0, 0, 1, 1, 0, 1}
	totalBarangRusak := 0

	for i := 0; i < len(barang); i++ {
		if barang[i] == 0 {
			totalBarangRusak++
		}
	}
	fmt.Println("Jumlah barang rusak: ", totalBarangRusak)
}
