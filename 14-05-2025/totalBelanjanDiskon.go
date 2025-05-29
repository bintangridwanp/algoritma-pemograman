package main

import "fmt"

func main() {

	hargaBarang := []int{120000, 75000, 200000, 90000}
	hasil := 0

	for i := 0; i < len(hargaBarang); i++ {
		if hargaBarang[i] > 100000 {
			hargaBarang[i] -= (hargaBarang[i] * 10) / 100
			hasil = hasil + hargaBarang[i]
		} else {
			hasil = hasil + hargaBarang[i]
		}
	}

	fmt.Println(hasil)

}
