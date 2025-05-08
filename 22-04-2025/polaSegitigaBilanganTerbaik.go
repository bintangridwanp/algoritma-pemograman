package main

import "fmt"

func main() {

	var input int

	fmt.Print("Enter an integer: ")
	fmt.Scanln(&input)

	for i := input; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}
