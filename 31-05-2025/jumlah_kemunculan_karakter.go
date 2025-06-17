package main

import "fmt"

func main() {
	var input string

	fmt.Print("Masukkan sebuah kata: ")
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		karakter := string(input[i])
		jumlahKemunculan := 0

		for j := 0; j < len(input); j++ {
			if karakter == string(input[j]) {
				jumlahKemunculan++
			}
		}

		fmt.Printf("Karakter '%s' muncul sebanyak %d kali\n", karakter, jumlahKemunculan)
	}

}
