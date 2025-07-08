package main

import "fmt"

func main() {

	nilai := [5][3]int{
		{80, 70, 90},
		{85, 88, 95},
		{60, 78, 84},
		{90, 65, 70},
		{70, 82, 89},
	}

	for pelajaran := 0; pelajaran < len(nilai[0]); pelajaran++ {
		nilaiTerbaik := nilai[0][pelajaran]
		for siswa := 1; siswa < len(nilai); siswa++ {
			if nilai[siswa][pelajaran] > nilaiTerbaik {
				nilaiTerbaik = nilai[siswa][pelajaran]
			}
		}
		fmt.Println("Nilai terbaik untuk mata pelajaran", pelajaran+1, "adalah", nilaiTerbaik)
	}

}
