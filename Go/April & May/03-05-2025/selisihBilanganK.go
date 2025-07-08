package main

import "fmt"

func main() {
	data := [...]int{1, 5, 3, 4, 2}
	k := 2

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if i != j {
				selisih := 0
				if data[i] > data[j] {
					selisih = data[i] - data[j]
				} else {
					selisih = data[j] - data[i]
				}

				if selisih == k {
					fmt.Println(data[i], "dan", data[j])
				}
			}
		}
	}
}
