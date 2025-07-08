package main

import "fmt"

func main() {

	terjual := []int{40, 60, 55, 45, 50, 48, 38}
	tidakTercapai, hasil := 0, 0
	target := 50

	for i := 0; i < len(terjual); i++ {
		if terjual[i] <= 50 {
			tidakTercapai++
			hasil += target - terjual[i]
		}
	}

	fmt.Println(tidakTercapai, "Hari target tidak tercapai", "dengan total kekurangan", hasil, "unit penjualan.")

}
