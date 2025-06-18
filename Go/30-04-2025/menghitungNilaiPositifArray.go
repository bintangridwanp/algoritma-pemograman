package main

import "fmt"

func main() {

	angka := [...]int{4, -2, 6, -8, 7, 10, -4}
	var hasil int

	for i := 0; i < len(angka); i++ {
		if angka[i] > 0 && angka[i]%2 == 0 {
			hasil = hasil + angka[i]
		}
	}
	fmt.Println(hasil)
}
