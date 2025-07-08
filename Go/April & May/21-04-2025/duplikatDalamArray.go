package main

import "fmt"

func main() {

	angka := [...]int{3, 5, 3, 2, 4, 5, 5}
	sudahDiCetak := [100]bool{}

	for i := 0; i < len(angka)-1; i++ {
		for j := i + 1; j < len(angka); j++ {
			if angka[i] == angka[j] && !sudahDiCetak[angka[i]] {
				fmt.Printf("%d ", angka[j])
				sudahDiCetak[angka[i]] = true
				break
			}
		}
	}
}
