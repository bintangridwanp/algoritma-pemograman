package main

import "fmt"

func main() {

	var angka int
	fmt.Print("Masukan bilangan biner: ")
	fmt.Scanln(&angka)

	biner := ""

	if angka < 0 {
		fmt.Println("Masukan bilangan biner tidak boleh negatif")
	} else if angka == 0 {
		fmt.Println("Bilangan desimal dari biner 0 adalah 0")
	} else {
		for angka > 0 {
			sisa := angka % 2
			biner = fmt.Sprintf("%d%s", sisa, biner)
			angka = angka / 2
		}
	}
	fmt.Println(biner)
}
