package main

import "fmt"

func main2() {

	angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 2, 5, 2, 3, 9, 0} 

	for i := 0; i < len(angka); i++{
		hasil := false
		for j := 0; j < i; j++ {
			if angka[i] == angka[j] {
				hasil = true
				break
			}
		}
		if hasil {
			fmt.Println("Angka ", angka[i], " muncul 2x")
		}
	}
}