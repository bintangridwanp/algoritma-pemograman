package main

import "fmt"

func main() {

	var input int
	fmt.Print("Enter seconds: ")
	fmt.Scanln(&input)
	Hours := input / 3600 // Menghitung jumlah jam
	minutes := input / 60 // Menghitung jumlah menit
	seconds := input % 60 // Menghitung sisa detik setelah dibagi menit
	fmt.Printf("%d detik = %d Hours %d menit dan %d detik\n", input, Hours, minutes, seconds)
}
