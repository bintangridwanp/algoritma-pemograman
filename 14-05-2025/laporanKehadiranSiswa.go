package main

import "fmt"

func main() {

	dataKehadiran := []rune{'H', 'A', 'H', 'A', 'H', 'H', 'A', 'H', 'A', 'H'}
	jumlahHadir := 0
	jumlahTidakHadir := 0

	for i := 1; i < len(dataKehadiran); i++ {
		if dataKehadiran[i] == 'H' {
			jumlahHadir++
		} else if dataKehadiran[i] == 'A' {
			jumlahTidakHadir++
		}
	}

	if jumlahTidakHadir > 3 {
		fmt.Println("Total tidak hadir: ", jumlahTidakHadir, "siswa perlu diambil tindakan")
	}

}
