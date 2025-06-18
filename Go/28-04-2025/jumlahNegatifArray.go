package main

import "fmt"

func main() {

	angka := [...]int{1, 5, -1, 9, -1, 2, 6, -1, 3}
	var hasil int

	for i := 0; i < len(angka); i++ {
		if angka[i] < 0 {
			hasil++
		}
	}

	fmt.Println(hasil)

}
