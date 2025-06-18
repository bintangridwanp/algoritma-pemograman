package main

import "fmt"

func main() {

	var hasil, input int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	for i := 0; i < input; i++ {
		if i%2 == 0 {
			hasil = hasil + 1
		}
	}

	fmt.Println("Total faktorial dari ", input, "is", hasil)
	
}
