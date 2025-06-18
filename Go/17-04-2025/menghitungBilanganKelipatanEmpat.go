package main

import "fmt"

func main() {

	var input, hasil int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		if i%4 == 0 {
			hasil = hasil + i
		}
	}

	fmt.Println("Hasil: ", hasil)

}
