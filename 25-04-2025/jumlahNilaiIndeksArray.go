package main

import "fmt"

func main() {

	var hasil int
	angka := [...]int{5, 10, 15, 20, 25, 30}

	for i := 0; i < len(angka); i++ {
		if i%2 != 0 {
			hasil = hasil + angka[i]
		}
	}

	fmt.Println(hasil)
}
