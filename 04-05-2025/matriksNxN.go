package main

import "fmt"

func main() {

	var n int
	fmt.Print("Masukkan ukuran matriks NxN: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || i+j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
