package main

import "fmt"

func main() {

	AbsensiSiswa := []string{"H", "A", "H", "S", "H", "A", "H"}
	hadir, alfa, sakit := 0, 0, 0

	for i := 0; i < len(AbsensiSiswa); i++ {
		if AbsensiSiswa[i] == "H" {
			hadir++
		} else if AbsensiSiswa[i] == "A" {
			alfa++
		} else if AbsensiSiswa[i] == "S" {
			sakit++
		}
	}
	fmt.Println("Hadir :", hadir, "Alfa :", alfa, "Sakit :", sakit)
}
