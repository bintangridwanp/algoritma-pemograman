package main

import "fmt"

func main() {

	nilaiUjian := []int{80, 70, 90, 60, 85}
	totalNilai := 0
	totalSiswa := 0

	for i := 0; i < len(nilaiUjian); i++ {
		totalNilai += nilaiUjian[i]
		totalSiswa++
		if totalNilai >= 75 {
			fmt.Println("siswa ke-", i+1, "lulus dengan nilai", nilaiUjian[i])
		}
	}

	hasil := float64(totalNilai) / float64(totalSiswa)
	fmt.Println("totalNilai", totalNilai, "hasil", hasil)

}
