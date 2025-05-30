package main

import "fmt"

func main() {

	nilaiTugas := []int{85, 90, 45, 60, 70, 40, 88, 92}
	tugasBaik := 0
	tugasPerluPerbaikan := 0

	for i := 0; i < len(nilaiTugas); i++ {
		if nilaiTugas[i] > 80 {
			tugasBaik++
		} else if nilaiTugas[i] < 50 {
			tugasPerluPerbaikan++
		}
	}

	fmt.Println("Tugas baik: ", tugasBaik)
	fmt.Println("Tugas perlu diperbaiki: ", tugasPerluPerbaikan)

}
