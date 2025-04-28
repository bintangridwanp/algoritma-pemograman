package main

import "fmt"

func main() {

	var input int;
	fmt.Print("Masukka sebuah angka: ");
	fmt.Scanln(&input);

	for input > 0 {
		hasil := input % 10;
		fmt.Println(hasil);
		input = input / 10;
	}

}