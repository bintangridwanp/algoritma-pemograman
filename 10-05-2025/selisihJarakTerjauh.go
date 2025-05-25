package main

import "fmt"

func main() {
	jarak := [...]int{120, 150, 200, 170, 100, 250, 130}

	// Inisialisasi nilai awal max dan min
	max := jarak[0]
	min := jarak[0]

	// Cari nilai maksimum dan minimum
	for i := 1; i < len(jarak); i++ {
		if jarak[i] > max {
			max = jarak[i]
		}
		if jarak[i] < min {
			min = jarak[i]
		}
	}

	// Hitung selisih
	selisih := max - min

	fmt.Println("Selisih terbesar:", selisih)
}
