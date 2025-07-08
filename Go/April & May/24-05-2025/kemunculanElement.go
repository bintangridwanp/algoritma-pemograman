package main

import "fmt"

func main() {
	angka := []int{1, 2, 2, 3, 1, 4, 2}

	for i := 0; i < len(angka); i++ {
		ada := false

		for j := 0; j < i; j++ {
			if angka[i] == angka[j] {
				ada = true
				break
			}
		}

		if ada {
			continue
		}

		jumlahKemunculan := 0
		for k := 0; k < len(angka); k++ {
			if angka[i] == angka[k] {
				jumlahKemunculan++
			}
		}
		fmt.Println("Angka", angka[i], "muncul", jumlahKemunculan, "kali")
	}
}
