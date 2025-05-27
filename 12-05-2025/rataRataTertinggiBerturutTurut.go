package main

import "fmt"

func main() {

	penjualan := []int{100, 200, 150, 300, 250, 400, 100, 200, 300, 100}

	rataTertinggi, hari, hasil := 0, 0, 0

	for i := 0; i < len(penjualan)-1; i++ {
		if penjualan[i] <= penjualan[i+1] {
			hasil += penjualan[i]
			hari++
			if i == len(penjualan)-2 {
				hasil += penjualan[i+1]
				hari++
			}
			if hari > 0 && hasil/hari > rataTertinggi {
				rataTertinggi = hasil / hari
				fmt.Println(rataTertinggi)
			}
		} else {
			hasil = 0
			hari = 0
		}
	}
	println(rataTertinggi)

}
