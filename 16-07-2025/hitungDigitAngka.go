package main

import "fmt"

func main2() {

	var input, hasil int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	if input == 0 {
		hasil = 1
	} else {
		for input > 0 {
			input = input / 10
			hasil++
		}
	}
	fmt.Println("The answer is", hasil)

}
