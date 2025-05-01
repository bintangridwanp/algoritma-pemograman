package main

import (
	"fmt"
)

func main1() {

	var hasil int
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	for input != 0 {
		hasil = input % 10
		fmt.Printf("%d", hasil)
		input /= 10
	}

}
