package main

import "fmt"

func main() {

	nilaiSiswa := []int{70, 80, 60, 75, 90, 85, 72, 88, 95, 68}
	siswaLulus := 0

	for i := 0; i < len(nilaiSiswa); i++ {
		if nilaiSiswa[i] >= 75 {
			siswaLulus++
		}
	}

	fmt.Println(siswaLulus, "Siswa lulus ujian dengan nilai >= 75")

}
