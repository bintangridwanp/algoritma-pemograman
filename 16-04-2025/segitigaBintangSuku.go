package main

import "fmt"

func main1() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	// Mulai dari sini kamu isi dengan loop-nya
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
