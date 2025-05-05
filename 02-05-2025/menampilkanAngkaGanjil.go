package main

import "fmt"

func main() {
	var input int
	fmt.Print("Input: ")
	fmt.Scan(&input)

	for i := input; i > 0; i-- {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}
