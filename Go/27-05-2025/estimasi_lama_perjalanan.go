package main

import "fmt"

func main() {

	var jarak, kecepatan, waktu int

	fmt.Print("Enter your jarak: ")
	fmt.Scan(&jarak)
	fmt.Print("Enter your kecepatan: ")
	fmt.Scan(&kecepatan)

	waktu = jarak / kecepatan // Menghitung waktu perjalanan dalam jam
	fmt.Printf("Lama perjalanan: %d jam\n", waktu)
}
