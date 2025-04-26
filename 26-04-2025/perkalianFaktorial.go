package main

import "fmt"

func main1() {
	var input int
	hasil := 1;
	fmt.Print("Masukan sebuah angka: ");
	fmt.Scanln(&input);

	for i := 1; i <= input; i++{
		fmt.Println("i", i);
		hasil = hasil * i;
		fmt.Println("hasil", hasil);
	}
	fmt.Println("Hasil perkalian faktorial dari", input, "adalah", hasil);

}