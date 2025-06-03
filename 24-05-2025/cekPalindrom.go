package main

import "fmt"

func main() {
	var input int
	fmt.Print("Enter a Number: ")
	fmt.Scanln(&input)

	original := input
	tempat := 0

	for input > 0 {
		digit := input % 10
		tempat = tempat*10 + digit
		input /= 10
	}

	fmt.Print()

	if original == tempat {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
