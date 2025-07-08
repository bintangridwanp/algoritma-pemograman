package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&angka)

	isPrima := true

	if angka <= 1 {
		isPrima = false
	} else {
		for i := 2; i <= angka/2; i++ {
			if angka%i == 0 {
				isPrima = false
				break
			}
		}
	}

	if isPrima {
		fmt.Println(angka, "adalah bilangan prima")
	} else {
		fmt.Println(angka, "bukan bilangan prima")
	}
}
