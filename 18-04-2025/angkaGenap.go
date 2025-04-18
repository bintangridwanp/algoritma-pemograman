package main

import "fmt"

func main () {

	angka := [...] int {10, 20, 30, 40, 50, 80, 25, 35, 100, 25, 50}

	x := 0

	for i := 0; i < len(angka); i++ {
		if angka[i] % 2 == 0 {
			x = x + 1
		}
	}

	fmt.Println("Jumlah angka genap adalah ", x);

}