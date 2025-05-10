package main

import "fmt"

func main() {

	pengunjung := [...]int{50, 80, 75, 90, 60, 85, 70}
	temp := pengunjung[0]
	hari := 0

	for i := 0; i < len(pengunjung); i++ {
		if pengunjung[i] > temp {
			temp = pengunjung[i]
			hari = i
		}
	}

	fmt.Println("Tertinggi adalah:", temp, "dihari ke:", hari)
}
