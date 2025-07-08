package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif:")
	fmt.Scanln(&input)

	jumlahGanjil := 0
	for i := 1; i <= input; i++ {
		if i%2 != 0 {
			jumlahGanjil += i
		}
	}

	fmt.Println("Jumlah bilangan ganjil dari 1 sampai", input, "adalah:", jumlahGanjil)
}
