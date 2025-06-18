package main

import "fmt"

func main() {
	var input, hasil int
	fmt.Print("Enter number: ")
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		hasil = i * input
	}

	fmt.Println("Here is the result:", hasil)
}
