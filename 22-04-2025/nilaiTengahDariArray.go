package main

import "fmt"

func main() {

	angka := []int{3, 1, 5, 4, 2}

	for i := 0; i < len(angka)-1; i++ {
		for j := 0; j < len(angka)-1-i; j++ {
			if angka[j] > angka[j+1] {
				temp := angka[j]
				angka[j] = angka[j+1]
				angka[j+1] = temp
			}
		}
	}

	hasil := angka[len(angka)/2]
	fmt.Println(hasil)
}
