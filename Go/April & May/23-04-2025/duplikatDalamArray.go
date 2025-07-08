package main

import "fmt"

func main() {

	angka := [...]int{3, 5, 3, 2, 4, 5, 5}
	var sudahDicetak [100]bool

	for i := 0; i < len(angka); i++ {
		for j := i + 1; j < len(angka); j++ {
			if angka[i] == angka[j] && !sudahDicetak[i] {
				fmt.Println("duplikat: ", angka[i])
				sudahDicetak[angka[i]] = true
				break
			}
		}
	}

}
