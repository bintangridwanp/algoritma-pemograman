package main

import "fmt"

func main() {

	var input int
	fmt.Println("Welcome to diskon belanja!")
	fmt.Scanln(&input)

	if input >= 100000 {
		diskon := input * 10 / 100 // Menghitung diskon 10%
		totalBayar := input - diskon
		fmt.Printf("Total belanja Anda adalah %d, diskon yang didapat adalah %d, total yang harus dibayar adalah %d\n", input, diskon, totalBayar)
	} else {
		fmt.Printf("Total belanja Anda adalah %d, tidak ada diskon yang diberikan.\n", input)
	}
}
