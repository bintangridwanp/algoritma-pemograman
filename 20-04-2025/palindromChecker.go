package main

import "fmt"

func main1() {

	var input string;

	fmt.Print("Masukan Kalimat: ")
	fmt.Scanln(&input)

	for i:=  0 ; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1]{
			fmt.Println("Not a palindrome")
			return
		}
	}
	fmt.Println("Palindrome")
}	