package main

import "fmt"

func main() {

	waktuKerja := []int{5, 7, 6, 4, 8, 3, 6}
	a, total := 0, 0

	for i := 0; i < len(waktuKerja); i++ {
		if waktuKerja[i] >= 6 {
			a++
			total += waktuKerja[i]
		}
	}

	fmt.Println("Hari produktif:", a, "Total jam produktif:", total)

}
