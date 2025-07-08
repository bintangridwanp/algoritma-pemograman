package main

import "fmt"

func main() {

	angka := []int{10, 3, 8, 2}
	var hasil1, hasil2 int

	for i := 0; i < len(angka)-1; i++ {
		a := angka[i]
		b := angka[i+1]

		if a > b {
			hasil1 = hasil1 + a - b
		} else if a < b {
			hasil2 = hasil2 + b - a
		}
	}
	fmt.Println(hasil1 + hasil2)
}
