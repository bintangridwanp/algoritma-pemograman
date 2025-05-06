package main

import "fmt"

func main() {

	var input int
	fmt.Print("Enter number: ")
	fmt.Scanln(&input)

	for i := input; i > 0; i-- {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}

}
