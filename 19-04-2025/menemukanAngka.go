package main

import "fmt"

func main() {
	angka := [...]int{10, 20, 30, 40, 50, 10, 20, 30, 40}

	for i := 0; i < len(angka); i++ {
		count := 0
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}
		if count == 1 {
			fmt.Println("Angka yang hanya muncul satu kali adalah:", angka[i])
			break
		}
	}
}
