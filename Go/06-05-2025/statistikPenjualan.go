package main

import "fmt"

func main() {

	produkA := []int{100, 200, 150, 300, 250}
	produkB := []int{80, 120, 90, 200, 160}
	produkC := []int{50, 70, 60, 0, 0}
	produkD := []int{50, 60, 70, 60, 80}

	totalPenjualanA := menghitungPenjualan(produkA)
	totalPenjualanB := menghitungPenjualan(produkB)
	totalPenjualanC := menghitungPenjualan(produkC)
	totalPenjualanD := menghitungPenjualan(produkD)

	totalPenjualan := []int{
		totalPenjualanA,
		totalPenjualanB,
		totalPenjualanC,
		totalPenjualanD,
	}

	temp := totalPenjualan[0]
	produkTerendah := 0

	for i := 0; i < len(totalPenjualan); i++ {
		if totalPenjualan[i] < temp {
			temp = totalPenjualan[i]
			produkTerendah = i
		}
	}

	fmt.Printf("Produk dengan penjualan terendah adalah Produk %c dengan total penjualan %d\n", 'A'+produkTerendah, temp)

}
func menghitungPenjualan(produk []int) int {

	var total int
	for i := 0; i < len(produk); i++ {
		total += produk[i]
	}
	return total

}
