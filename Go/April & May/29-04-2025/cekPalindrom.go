package main

import "fmt"

func main1() {
	var kata string

	fmt.Print("Masukan sebuah kata: ")
	fmt.Scanln(&kata)

	var hasil bool
	hasil = true

	for i := 0; i < len(kata); i++ {
		if kata[i] != kata[len(kata)-i-1] {
			hasil = false
		}
	}

	if hasil == true {
		fmt.Println("Polindrom")
	} else {
		fmt.Println(" Not Polindrom")
	}

}
