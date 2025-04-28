package main

import "fmt"

func main2() {

	var input int
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&input);

	for i := 1; i <= input; i++ {

		if i % 2 == 0 {
			fmt.Println(i, "Genap");
		} else {
			fmt.Println(i, "Ganjil");
		}
	}
}