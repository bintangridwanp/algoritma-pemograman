package main

import "fmt"

func main1() {
	var input int
	fmt.Print("Input: ")
	fmt.Scan(&input)

	original := input
	reversed := 0

	for input > 0 {
		digit := input % 10
		reversed = reversed*10 + digit
		input = input / 10
	}

	if original == reversed {
		fmt.Printf("%d adalah palindrom\n", original)
	} else {
		fmt.Printf("%d bukan palindrom\n", original)
	}
}
