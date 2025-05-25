package main

import "fmt"

func main() {

	totalBelanja := []int{50000, 75000, 60000, 80000, 100000}
	total := 0
	a := 0

	for i := 0; i < len(totalBelanja); i++ {
		total += totalBelanja[i]
		a++
	}

	hasil := total / a

	fmt.Println("Total: ", total, "Rata-rata: ", hasil)
}
