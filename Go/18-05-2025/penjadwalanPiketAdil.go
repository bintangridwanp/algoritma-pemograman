package main

import "fmt"

func main() {
	jumlahHari := 100
	jumlahSiswa := 4

	for i := 1; i <= jumlahHari; i++ {
		siswa := (i-1)%jumlahSiswa + 1
		fmt.Printf("Hari %d: Siswa %d\n", i, siswa)
	}
}
