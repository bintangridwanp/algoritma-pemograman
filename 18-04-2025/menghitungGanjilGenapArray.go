package main

import "fmt"

func main() {

	angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var ganjil, genap int

	for i := 0; i < len(angka); i++ {
		if angka[i]%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}

	fmt.Println("Total angka ganjil : ", ganjil)
	fmt.Println("Total angka ganjil : ", ganjil)
}
