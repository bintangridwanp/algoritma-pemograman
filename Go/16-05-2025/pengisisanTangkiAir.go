package main

import "fmt"

func main() {

	jumlahVolume := []int{100, 200, 150, 300, 180, 100}
	jam := 0

	for i := 0; i < len(jumlahVolume); i++ {
		if jumlahVolume[i] > 0 {
			jam += jumlahVolume[i]
			if jam > 1000 {
				fmt.Println("Tangki air penuh pada jam ke", i)
			}
		}
	}

}
