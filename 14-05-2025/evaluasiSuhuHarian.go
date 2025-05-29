package main

import "fmt"

func main() {

	suhuHarian := []int{22, 18, 19, 25, 23, 17, 21}
	jumlahSuhu := 0

	for i := 1; i < len(suhuHarian); i++ {
		if suhuHarian[i] > 20 {
			jumlahSuhu++
		}
	}

	if jumlahSuhu > 2 {
		fmt.Println("Hari berisiko", jumlahSuhu, "Tidak aman")
	} else {
		fmt.Println("Aman")
	}

}
