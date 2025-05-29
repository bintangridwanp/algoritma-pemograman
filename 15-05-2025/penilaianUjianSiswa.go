package main

import "fmt"

func main() {

	nilaiSsiwa := []int{80, 60, 90, 72, 77, 40, 85}
	lulus := 0
	tidakLulus := 0

	for i := 0; i < len(nilaiSsiwa); i++ {
		if nilaiSsiwa[i] >= 75 {
			lulus++
		} else {
			tidakLulus++
		}
	}

	fmt.Println("Jumlah siswa yang lulus:", lulus, "siswa")
	fmt.Println("Jumlah siswa yang tidak lulus", tidakLulus, "siswa")
}
