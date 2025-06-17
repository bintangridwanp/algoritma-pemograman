package main

import "fmt"

func main() {
	angka := []int{7, 3, 5, 3, 1}
	terkecil := angka[0]
	terkecilKedua := angka[0]

	for i := 0; i < len(angka); i++ {
		if angka[i] < terkecil {
			terkecilKedua = terkecil
			terkecil = angka[i]
		} else if angka[i] > terkecil && angka[i] < terkecilKedua {
			terkecilKedua = angka[i]
		}
	}

	fmt.Println("Angka terkecil kedua adalah", terkecilKedua)
}
