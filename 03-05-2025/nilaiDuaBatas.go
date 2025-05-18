package main

import "fmt"

func main() {

	angka := []int{2, 4, 7, 10, 13, 16}

	a := 5
	b := 13
	hasil := 0

	for i := 0; i < len(angka); i++ {
		if angka[i] >= a && angka[i] <= b {
			print(angka[i], " ")
			hasil++
		}
	}

	fmt.Println("Total = ", hasil, "yaitu dari", a, "sampai", b)

}
