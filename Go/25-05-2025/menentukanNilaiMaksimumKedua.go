package main

import "fmt"

func main() {
	angka := []int{3, 8, 2, 10, 6, 10}
	nilaiMaksimum := angka[0]
	nilaiMaksimumKedua := angka[0]

	for i := 0; i < len(angka); i++ {
		if angka[i] > nilaiMaksimum {
			nilaiMaksimumKedua = nilaiMaksimum
			nilaiMaksimum = angka[i]
		} else if angka[i] > nilaiMaksimumKedua && angka[i] < nilaiMaksimum {
			nilaiMaksimumKedua = angka[i]
		}
	}
	fmt.Println("Nilai maksimum kedua adalah:", nilaiMaksimumKedua)
}
