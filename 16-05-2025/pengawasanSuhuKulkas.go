package main

import "fmt"

func main() {

	suhuKulkas := []int{4, 3, 6, 5, 7, 3, 2, 8}
	tidakAman := 0

	for i := 0; i < len(suhuKulkas); i++ {
		if suhuKulkas[i] > 5 {
			tidakAman++
		}
	}

	fmt.Println("Suhu kulkas tidak aman sebanyak: ", tidakAman)

}
