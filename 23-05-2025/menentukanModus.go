package main

import "fmt"

func main() {
	angka := []int{7, 1, 3, 1, 7, 7, 3, 3, 3}

	maksFrekuensi := 0
	modus := angka[0]

	for i := 0; i < len(angka); i++ {
		count := 1

		for j := i + 1; j < len(angka); j++ {
			if angka[i] == angka[j] {
				count++
			}
		}

		if count > maksFrekuensi {
			maksFrekuensi = count
			modus = angka[i]
		}
	}

	fmt.Println("Modus:", modus)
}
