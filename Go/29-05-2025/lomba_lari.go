package main

import "fmt"

func main() {

	kecepatanLari := []int{15, 13, 18, 11, 14}
	tercepat := kecepatanLari[0]

	for i := 0; i < len(kecepatanLari); i++ {
		if kecepatanLari[i] < tercepat {
			tercepat = kecepatanLari[i]
		}
	}

	fmt.Println("Kecepatan lari tercepat adalah", tercepat, "km/jam")
}
