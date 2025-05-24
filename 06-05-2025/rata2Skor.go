package main

import "fmt"

func main() {

	Tim1 := []int{80, 90, 85, 70}
	Tim2 := []int{60, 75, 80, 95}
	Tim3 := []int{90, 85, 88, 92}

	rataRataTim1 := hitungRataRata(Tim1)
	rataRataTim2 := hitungRataRata(Tim2)
	rataRataTim3 := hitungRataRata(Tim3)

	if rataRataTim1 > rataRataTim2 && rataRataTim1 > rataRataTim3 {
		fmt.Println("Tim 1 adalah tim dengan rata-rata skor tertinggi:", rataRataTim1)
	} else if rataRataTim2 > rataRataTim1 && rataRataTim2 > rataRataTim3 {
		fmt.Println("Tim 2 adalah tim dengan rata-rata skor tertinggi:", rataRataTim2)
	} else if rataRataTim3 > rataRataTim1 && rataRataTim3 > rataRataTim2 {
		fmt.Println("Tim 3 adalah tim dengan rata-rata skor tertinggi:", rataRataTim3)
	} else {
		fmt.Println("Ada dua atau lebih tim dengan rata-rata skor tertinggi yang sama.")
	}

}

func hitungRataRata(nilai []int) float64 {
	var total int
	for i := 0; i < len(nilai); i++ {
		total += nilai[i]
	}

	return float64(total) / float64(len(nilai))
}
