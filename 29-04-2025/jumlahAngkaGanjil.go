package main

import "fmt"

func main() {

	var input int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	var hasil int
	for i := 0; i <= input; i++ {
		if i%2 != 0 {
			hasil = hasil + i
		}
	}

	fmt.Println(hasil)

}
