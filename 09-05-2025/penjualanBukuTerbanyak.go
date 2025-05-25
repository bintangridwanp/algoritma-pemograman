package main

import "fmt"

func main() {

	penjualanBuku := []int{30, 45, 22, 60, 55}
	terbanyak := penjualanBuku[0]
	a := 0

	for i := 1; i < len(penjualanBuku); i++ {
		if penjualanBuku[i] > terbanyak {
			terbanyak = penjualanBuku[i]
			a = i
		}
	}
	fmt.Println("Jumlah penjualan buku terbanyak adalah:", terbanyak, "pada indeks ke-", a)
}
