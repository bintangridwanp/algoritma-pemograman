package main

import "fmt"

func main2() {

	var max, min int

	angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max = angka[0]
	min = angka[0]

	for i := 0; i < len(angka); i++ {
		if angka[i] > max {
			max = angka[i]
		}
		if angka[i] < min {
			min = angka[i]
		}
	}

	fmt.Println("Nilai max dari array angka adalah: ", max)
	fmt.Println("Nilai min dari array angka adalah: ", min)

}
