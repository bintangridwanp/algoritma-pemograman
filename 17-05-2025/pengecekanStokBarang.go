package main

import "fmt"

func main() {

	stokBarang := []int{12, 8, 15, 9, 11, 6, 14}
	totalRestock := 0

	for i := 0; i < len(stokBarang); i++ {
		if stokBarang[i] < 10 {
			totalRestock++
		}
	}

	fmt.Println("Toko harus restok sebanyak", totalRestock, "kali.")
}
