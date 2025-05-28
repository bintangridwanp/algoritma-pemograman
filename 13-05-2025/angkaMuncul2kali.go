package main

import "fmt"

func main() {
	angka := []int{1, 2, 3, 2, 4, 5, 1, 6, 7, 5, 5}

	for i := 0; i < len(angka); i++ {
		count := 1
		for j := i + 1; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}
		if count == 2 {
			fmt.Printf("Angka %d muncul 2 kali\n", angka[i])
		}
	}
}
