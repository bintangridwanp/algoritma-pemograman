package main

import "fmt"

func main() {

	var hasil int = 0

	angka := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(angka); i++ {
		hasil += angka[i] 
	}
	fmt.Println("Hasil Penjumlahan Array adalah ", hasil)
}