package main

import "fmt"

func main() {

	totalPengunjung := []int{101, 102, 103, 101, 104, 102, 105, 106}

	for i := 0; i < len(totalPengunjung); i++ {
		for j := i + 1; j < len(totalPengunjung); j++ {
			if totalPengunjung[i] == totalPengunjung[j] {
				fmt.Println("Duplikat ID:", totalPengunjung[i])
				break
			}
		}
	}

}
