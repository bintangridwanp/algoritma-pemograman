package main

import "fmt"

func main() {

	gaji := []int{100000, 100000, 100000, 120000, 100000, 100000, 100000}
	total := 0

	for i := 0; i < len(gaji); i++ {
		total += gaji[i]
	}

	fmt.Println("Total gaji karyawan adalah", total, "rupiah")
}
