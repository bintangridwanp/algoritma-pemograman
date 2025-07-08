package main

import "fmt"

func main() {

	pesanan := []int{101, 102, 103, 104, 105, 102, 106, 103, 107, 108}
	a := []int{}

	for i := 0; i < len(pesanan); i++ {
		for j := i + 1; j < len(pesanan); j++ {
			if pesanan[i] == pesanan[j] {
				a = append(a, pesanan[i])
			}
		}
	}

	if len(a) == 0 {
		fmt.Println("Nothing found")
	} else {
		for i := 0; i < len(a); i++ {
			fmt.Println("Duplikat pesanan ditemukan:", a[i])
		}
	}

}
