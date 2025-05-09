package main

import "fmt"

func main() {

	var input int
	fmt.Print("Enter an number: ")
	fmt.Scanln(&input)

	var hasil, n int

	for input != 0 {
		hasil = input % 10
		if hasil%2 == 0 {
			n++
		}
		input /= 10
	}

	fmt.Println(n)

}
