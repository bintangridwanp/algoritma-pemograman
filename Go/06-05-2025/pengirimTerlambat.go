package main

import "fmt"

func main() {

	waktuPengiriman := []int{30, 50, 40, 60, 35, 55, 45}
	var totalTerlambat int

	for i := 0; i < len(waktuPengiriman); i++ {
		if waktuPengiriman[i] > 45 {
			fmt.Println("Pengiriman ke-", i+1, "terlambat dengan waktu:", waktuPengiriman[i], "menit")
			totalTerlambat++
		}
	}

	totalPengiriman := len(waktuPengiriman)
	persenTerlambat := (float64(totalTerlambat) / float64(totalPengiriman)) * 100

	fmt.Println("Total pengiriman terlambat:", totalTerlambat)
	fmt.Printf("Persentase pengiriman terlambat: %.2f%%\n", persenTerlambat)
}
