package main

import "fmt"

func main1() {

	var input int;

	fmt.Print("Masukan sebuah angka: ")
	fmt.Scanln(&input)

	if input % 2 == 0 {
		fmt.Println(input, " adalah sebuah bilangan genap");
	} else{
		fmt.Println(input,  " adalah sebuah bilangan ganjil");
	}
}