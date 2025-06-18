package main

import "fmt"

func main() {
	var jarak float64
	var kecepatan float64

	fmt.Print("Masukkan jarak (km): ")
	fmt.Scanln(&jarak)

	fmt.Print("Masukkan kecepatan (km/jam): ")
	fmt.Scanln(&kecepatan)

	if kecepatan == 0 {
		fmt.Println("Kecepatan tidak boleh nol.")
		return
	}

	waktu := jarak / kecepatan

	if waktu < 1 {
		menit := waktu * 60
		fmt.Printf("Waktu tempuh: %.0f menit\n", menit)
	} else {
		fmt.Printf("Waktu tempuh: %.2f jam\n", waktu)
	}
}
