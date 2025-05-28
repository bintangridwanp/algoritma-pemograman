package main

import "fmt"

func main() {

	angka := []int{3, 0, 4, 5, -1, 0, 2, 0, -3}
	jumlahPositif := 0

	for i := 0; i < len(angka); i++ {
		if angka[i] == 0 && i < len(angka)-1 {
			if angka[i+1] > 0 {
				fmt.Println(angka[i+1])
				jumlahPositif++
			} else {
				fmt.Println("Tidak ada angka positif setelah 0")
			}
		}
	}

	fmt.Println("Jumlah angka positif setelah 0:", jumlahPositif, "angka positif ditemukan")

}
