package main

import "fmt"

func main() {

	nilai := [...]int{80, 60, 75, 90, 55, 70, 88, 45, 77, 95}
	var lulus, tidakLulus int

	for i := 0; i < len(nilai); i++ {
		if nilai[i] >= 75 {
			lulus++
		} else {
			tidakLulus++
		}
	}
	fmt.Println("Lulus:", lulus)
	fmt.Println("TidakLulus:", tidakLulus)
}
