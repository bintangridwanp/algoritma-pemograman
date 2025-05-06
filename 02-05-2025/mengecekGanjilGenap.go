package main

import "fmt"

func main() {

	var a, b int
	angka := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 1; i <= len(angka); i++ {
		if i%2 == 0 {
			a = a + i
		} else {
			b = b + i
		}
	}
	fmt.Println(a)
	fmt.Println(b)
}
