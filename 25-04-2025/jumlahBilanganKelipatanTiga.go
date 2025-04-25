package main

import "fmt"

func main() {

	var input int;
	hasil := 0

	fmt.Print("Masukan sebuah angka: ");
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		if i % 3 == 0 {
			hasil = hasil + i
		}
	}

	fmt.Println("Jumlah bilangan kelipatan 3 dari 1 sampai", input, "adalah", hasil);


}