package main

import "fmt"

func main() {

	angka := [...]int{1, 2, 3, 4, 5}
	var total int
	total = 6

	for i := 0; i < len(angka); i++ {
		for j := i + 1; j < len(angka); j++ {
			if angka[i]+angka[j] == total {
				fmt.Println(angka[i], "+", angka[j])
			}
		}
	}
}
