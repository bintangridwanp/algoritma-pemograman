package main

import "fmt"

func main() {

	penjualan := [...]int{98, 120, 101, 99, 105, 80, 110}
	var target int

	for i := 0; i < len(penjualan); i++ {
		if penjualan[i] > 100 {
			target++
		}
	}
	fmt.Println(target, "Melebihi Target")
}
