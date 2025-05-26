package main

import "fmt"

func main() {

	tabungang := []int{500, 700, 10, 10, 10, 0, 800, 900, 1000, 1100, 0, 200, 300, 400, 500}
	a, tempat := 0, 0

	for i := 0; i < len(tabungang); i++ {
		if tabungang[i] != 0 {
			a++
			if a > tempat {
				tempat = a
			}
		} else if tabungang[i] == 0 {
			a = 0
		}
	}

	fmt.Println(tempat)

}
