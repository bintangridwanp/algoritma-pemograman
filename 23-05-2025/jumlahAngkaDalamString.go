package main

import "fmt"

func main() {

	var input string
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	jumlahAngka := 0

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			jumlahAngka = jumlahAngka + int(input[i]-'0')
		}
	}

	fmt.Println(jumlahAngka)

}
