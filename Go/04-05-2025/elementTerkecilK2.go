package main

import "fmt"

func main() {
	data := [...]int{8, 3, 7, 1, 4}

	min1 := data[0]
	min2 := 9999999

	for i := 1; i < len(data); i++ {
		if data[i] < min1 {
			min1 = data[i]
		}
	}

	for i := 0; i < len(data); i++ {
		if data[i] > min1 && data[i] < min2 {
			min2 = data[i]
		}
	}

	fmt.Println("Nilai terkecil ke-2 adalah:", min2)
}
