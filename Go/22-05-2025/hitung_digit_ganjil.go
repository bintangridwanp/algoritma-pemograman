package main

import "fmt"

func main() {

	var angka int
	totalGanjil := 0

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scanln(&angka)

	for angka > 0 {
		digit := angka % 10 // Mengambil digit terakhir
		if digit%2 != 0 {   // Mengecek apakah digit tersebut ganjil
			totalGanjil++ // Menambah jumlah digit ganjil
		}
		angka /= 10 // Menghapus digit terakhir
	}
	fmt.Println("Jumlah digit ganjil adalah: ", totalGanjil)

}
