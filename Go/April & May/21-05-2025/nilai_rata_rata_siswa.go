package main

import "fmt"

func main() {

	nilai := []int{80, 90, 75, 85, 95}
	jumlahNilai, jumlahSiswa := 0, 0
	nilaiTerendah := nilai[0]
	nilaiTertingi := nilai[0]
	siswaTertinggi, siswaTerendah := 0, 0

	for i := 0; i < len(nilai); i++ {
		jumlahNilai += nilai[i]
		jumlahSiswa++
		if nilaiTertingi < nilai[i]+1 {
			nilaiTertingi = nilai[i]
			siswaTertinggi = i
		}
		if nilaiTerendah > nilai[i]+1 {
			nilaiTerendah = nilai[i]
			siswaTerendah = i
		}

	}
	nilaiRataRata := float64(jumlahNilai) / float64(jumlahSiswa)
	fmt.Println("Nilai rata-rata:", nilaiRataRata)
	fmt.Println("Jumlah siswa:", jumlahSiswa)
	fmt.Println("Jumlah nilai:", jumlahNilai)
	fmt.Println("Nilai terendah:", nilaiTerendah, "pada siswa", siswaTerendah)
	fmt.Println("Nilai tertinggi:", nilaiTertingi, "pada siswa", siswaTertinggi)
}
