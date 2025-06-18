package main

import "fmt"

func main() {

	var input int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		hasil := i * i
		fmt.Println(hasil)
	}

}
