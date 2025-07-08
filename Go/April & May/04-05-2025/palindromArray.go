package main

import "fmt"

func main() {

	angka := []int{1, 2, 3, 2, 2}
	palindrome := true

	for i := 0; i < len(angka)/2; i++ {
		if angka[i] != angka[len(angka)-1-i] {
			palindrome = false
			break
		}
	}

	if palindrome {
		fmt.Println("Array adalah palindrome")
	} else {
		fmt.Println("Array bukan palindrome")
	}

}
