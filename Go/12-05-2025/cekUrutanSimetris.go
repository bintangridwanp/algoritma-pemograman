package main

import "fmt"

func main() {

	angka := []int{1, 3, 5, 7, 5, 3, 1}
	//angka := []int{2, 4, 6, 6, 4, 2}

	//for i := 0; i < len(angka); i++ {
	//	fmt.Println(angka[len(angka)-i-1])
	//}

	simetris := true

	for i := 0; i < len(angka)/2; i++ {
		fmt.Println(angka[len(angka)-1-i])
		if angka[i] != angka[len(angka)-1-i] {
			simetris = false
			break
		}
	}

	if simetris {
		println("Urutan angka simetris")
	} else {
		println("Urutan angka tidak simetris")
	}
}
