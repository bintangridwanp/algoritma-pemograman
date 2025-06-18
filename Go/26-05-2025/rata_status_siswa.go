package main

func main() {

	nilaiSiswa := []int{80, 90, 60, 75, 85}
	totalNilai, totalSiswa := 0, 0

	for i := 0; i < len(nilaiSiswa); i++ {
		totalNilai += nilaiSiswa[i]
		totalSiswa++
	}
	rataRata := totalNilai / totalSiswa

	if rataRata >= 75 {
		println("Rata-rata siswa lulus dengan nilai", rataRata)
	} else {
		println("Rata-rata siswa tidak lulus dengan nilai", rataRata)
	}
}
