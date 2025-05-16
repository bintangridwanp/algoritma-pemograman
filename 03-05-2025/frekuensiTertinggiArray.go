package main

import "fmt"

func main() {
	data := [...]int{1, 2, 2, 3, 3, 3, 4, 2, 3}

	var frekuensi [101]int

	for i := 0; i < len(data); i++ {
		frekuensi[data[i]] = frekuensi[data[i]] + 1
		fmt.Println(frekuensi[data[i]])
	}

	maks := 0
	for i := 0; i < len(frekuensi); i++ {
		if frekuensi[i] > maks {
			maks = frekuensi[i]
		}
	}

	fmt.Println("Nilai paling sering muncul:")
	for i := 0; i < len(frekuensi); i++ {
		if frekuensi[i] == maks {
			fmt.Printf("%d (muncul %d kali)\n", i, maks)
		}
	}
}
