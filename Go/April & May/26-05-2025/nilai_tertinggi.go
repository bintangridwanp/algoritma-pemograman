package main

import "fmt"

func main() {

	angka := []int{78, 92, 85, 74, 90, 88, 95}
	tertinggi := angka[0]

	for i := 0; i < len(angka); i++ {
		if angka[i] > tertinggi {
			tertinggi = angka[i]
		}
	}
	fmt.Println("Nilai tertinggi adalah", tertinggi)
}
