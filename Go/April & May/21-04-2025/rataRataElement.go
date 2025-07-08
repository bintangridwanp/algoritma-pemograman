package main

import "fmt"

func main() {

	angka := [...]int{4, 6, 8, 10, 11}
	var hasil float32

	for i := 0; i < len(angka); i++ {
		hasil += float32(angka[i])
	}

	temp := (hasil / float32(len(angka)))

	hasil2 := 0

	for i := 0; i < len(angka); i++ {
		if float32(angka[i]) > temp {
			hasil2++
		}
	}

	fmt.Println(hasil2)
}
