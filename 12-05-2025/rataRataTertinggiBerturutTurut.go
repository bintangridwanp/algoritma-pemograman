package main

import (
	"fmt"
)

func main() {
	data := []int{100, 200, 150, 300, 250, 400, 100, 200, 300, 100}

	maxAvg := 0.0
	var maxSegment [3]int

	for i := 0; i <= len(data)-3; i++ {
		sum := data[i] + data[i+1] + data[i+2]
		avg := float64(sum) / 3

		if avg > maxAvg {
			maxAvg = avg
			maxSegment = [3]int{data[i], data[i+1], data[i+2]}
		}
	}

	fmt.Printf("Rata-rata tertinggi: %.2f (dari %v)\n", maxAvg, maxSegment)
}
