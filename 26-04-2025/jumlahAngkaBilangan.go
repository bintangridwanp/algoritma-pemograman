package main

import "fmt"

func main2() {

	var input, hasil int;
	fmt.Print("Masukan sebuah angka: " );
	fmt.Scanln(&input);

	for input > 0{
		hasil = hasil + input % 10;
		fmt.Println("hasil", hasil);
		fmt.Println("input", input);
		input = input / 10;
	}

	fmt.Println(hasil)
} 