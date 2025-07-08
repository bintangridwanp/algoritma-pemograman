package main

import "fmt"

func main() {

	suhu := []int{28, 36, 19, 22, 35, 38, 18}
	jumlahSuhuEkstrem := 0

	for i := 0; i < len(suhu); i++ {
		if suhu[i] > 35 || suhu[i] < 20 {
			jumlahSuhuEkstrem++
		}
	}
	fmt.Println(jumlahSuhuEkstrem, "hari ekstrem")
}
