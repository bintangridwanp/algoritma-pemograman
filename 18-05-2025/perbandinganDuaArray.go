package main

import "fmt"

func main() {

	gudang1 := []int{101, 102, 103, 104, 105}
	gudang2 := []int{102, 104, 106}

	tempat1 := []int{}
	tempat2 := []int{}

	for i := 0; i < len(gudang1); i++ {
		found := false
		for j := 0; j < len(gudang2); j++ {
			if gudang1[i] == gudang2[j] {
				found = true
				break
			}
		}
		if found {
			tempat1 = append(tempat1, gudang1[i])
		} else {
			tempat2 = append(tempat2, gudang1[i])
		}
	}

	fmt.Println(tempat1)
	fmt.Println(tempat2)

}
