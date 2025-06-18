package main

import "fmt"

func main() {

	angka := []int{1, 2, 3, 4, 5, 6}
	ganjil := 0
	genap := 0

	for i := 0; i < len(angka); i++ {
		if angka[i]%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}

	fmt.Println("Jumlah genap: ", genap)
	fmt.Println("Jumlah ganjil: ", ganjil)

}
