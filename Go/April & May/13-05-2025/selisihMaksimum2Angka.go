package main

import "fmt"

func main() {
	angka := []int{7, 1, 5, 3, 6, 4}

	if len(angka) < 2 {
		fmt.Println("Array harus memiliki minimal dua elemen.")
		return
	}

	minSoFar := angka[0]
	selisihMaks := 0

	for i := 1; i < len(angka); i++ {
		selisih := angka[i] - minSoFar
		if selisih > selisihMaks {
			selisihMaks = selisih
		}
		if angka[i] < minSoFar {
			minSoFar = angka[i]
		}
	}

	fmt.Println("Selisih terbesar:", selisihMaks)
}
