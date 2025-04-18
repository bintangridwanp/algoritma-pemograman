package main

import "fmt"

func main1(){

	angka := [...] int {10, 20, 30, 40, 50, 80, 25, 35, 100, 25, 50}

	temporary := 0;

	for i := 0; i < len(angka); i++ {
		if angka[i] > temporary {
			temporary = angka[i]
		}
	}

	fmt.Println("Nilai maksimum adalah ", temporary);
}