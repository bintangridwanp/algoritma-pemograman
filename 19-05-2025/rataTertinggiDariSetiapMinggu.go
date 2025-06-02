package main

import (
	"fmt"
)

func main() {

	nilai := []float64{
		80, 70, 75, 85, 90, 78, 60, // minggu 1
		88, 92, 77, 85, 82, 80, 79, // minggu 2
		65, 70, 75, 72, 68, 74, 71, // minggu 3
		90, 95, 93, 89, 91, 88, 87, // minggu 4
		70, 70,
	}

	jumlahHari := len(nilai)
	minggu := 1
	tertinggi := 0.0
	mingguTerbaik := 1

	for i := 0; i < jumlahHari; i += 7 {
		akhir := i + 7
		if akhir > jumlahHari {
			akhir = jumlahHari
		}
		total := 0.0
		for j := i; j < akhir; j++ {
			total += nilai[j]
		}
		rata := total / float64(akhir-i)
		fmt.Printf("Rata-rata minggu %d: %.1f\n", minggu, rata)

		if rata > tertinggi {
			tertinggi = rata
			mingguTerbaik = minggu
		}

		minggu++
	}

	fmt.Printf("Minggu terbaik: minggu %d\n", mingguTerbaik)
}
