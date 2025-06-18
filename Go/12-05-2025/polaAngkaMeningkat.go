package main

import "fmt"

func main() {

	//angka := []int{10, 20, 30, 40, 50, 60}
	angka := []int{10, 20, 30, 28, 50, 60}

	polaMeningkat := true

	for i := 0; i < len(angka)-1; i++ {
		if angka[i] >= angka[i+1] {
			fmt.Println(i)
			polaMeningkat = false
			break
		}
		fmt.Println(i)
	}

	if polaMeningkat {
		println("Pola angka meningkat")
	} else {
		println("Pola angka tidak meningkat")
	}

}
