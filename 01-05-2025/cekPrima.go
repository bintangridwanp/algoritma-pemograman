package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	isPrime := true

	if n < 2 { // 3 < 2 true
		isPrime = false
	} else {
		for i := 2; i < n; i++ { // 1 = 2, 2 < 3
			if n%i == 0 { // 3 mod 2 == 0
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(n, "adalah bilangan prima")
	} else {
		fmt.Println(n, "bukan bilangan prima")
	}
}
