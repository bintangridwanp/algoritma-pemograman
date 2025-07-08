package main

import "fmt"

func main() {
	nilai := []int{80, 65, 90, 45, 77, 70}
	Lulus := 0
	TidakLulus := 0

	for i := 0; i < len(nilai); i++ {
		if nilai[i] >= 75 {
			Lulus++
		} else if nilai[i] <= 75 {
			TidakLulus++
		}
	}

	fmt.Print("Jumlah siswa yang lulus: ", Lulus, " siswa")
	fmt.Print("\nJumlah siswa yang tidak lulus: ", TidakLulus, " siswa")

}
