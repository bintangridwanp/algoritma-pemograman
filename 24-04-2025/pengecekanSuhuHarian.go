package main

import "fmt"

func main() {

	suhu := [...]int{28, 32, 31, 29, 30, 35, 33}
	var sum int

	for i := 0; i < len(suhu); i++ {
		if suhu[i] > 30 {
			sum++
		}
	}
	fmt.Println(sum, "Hari panas")
}
