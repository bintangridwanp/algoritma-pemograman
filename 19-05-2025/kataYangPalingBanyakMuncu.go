package main

import "fmt"

func main() {
	kata := []string{"belajar", "main", "belajar", "main", "belajar", "tidur", "main"}

	kataTerbanyak := ""
	jumlahTerbanyak := 0

	for i := 0; i < len(kata); i++ {
		count := 1
		sudahDihitung := false
		for k := 0; k < i; k++ {
			if kata[i] == kata[k] {
				sudahDihitung = true
				break
			}
		}
		if sudahDihitung {
			continue
		}

		for j := i + 1; j < len(kata); j++ {
			if kata[i] == kata[j] {
				count++
			}
		}
		if count > jumlahTerbanyak {
			jumlahTerbanyak = count
			kataTerbanyak = kata[i]
		}
	}

	fmt.Printf("Kata paling sering: %s (%d kali)\n", kataTerbanyak, jumlahTerbanyak)
}
