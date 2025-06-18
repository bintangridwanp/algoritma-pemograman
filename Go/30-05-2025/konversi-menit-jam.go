package main

import "fmt"

func main() {

	var menit, jam, sisaMenit int
	fmt.Print("Masukkan jumlah menit:")
	fmt.Scanln(&menit)
	jam = menit / 60       // Menghitung jumlah jam
	sisaMenit = menit % 60 // Menghitung sisa menit setelah dibagi jam
	fmt.Printf("%d menit = %d jam dan %d menit\n", menit, jam, sisaMenit)

}
