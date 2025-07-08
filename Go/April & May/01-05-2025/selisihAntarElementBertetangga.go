package main

import "fmt"

func main() {

	angka := [...]int{5, 1, 9, 3}
	var hasil int

	for i := 0; i < len(angka)-1; i++ {
		a := angka[i]
		b := angka[i+1]

		if a > b {
			hasil = hasil + a - b
		} else if b > a {
			hasil = hasil + b - a
		}
	}

	fmt.Println(hasil)

}
