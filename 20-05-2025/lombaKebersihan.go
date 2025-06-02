package main

import "fmt"

func main() {

	nilaiLomba := []int{80, 75, 90, 85, 70}
	nilaiTerendah := nilaiLomba[0]
	hasil, hariTerendah := 0, 1

	for i := 0; i < len(nilaiLomba); i++ {
		if nilaiLomba[i] < nilaiTerendah {
			nilaiTerendah = nilaiLomba[i]
			hariTerendah = i + 1
		}
		hasil = hasil + nilaiLomba[i]
	}
	// Menghitung rata-rata
	rataRata := hasil / len(nilaiLomba)

	fmt.Println("Nilai terendah:", nilaiTerendah)
	fmt.Println("rata rata nilai:", rataRata, "pada hari ke", hariTerendah)

}
