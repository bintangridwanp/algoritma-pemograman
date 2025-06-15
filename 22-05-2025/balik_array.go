package main

import "fmt"

func main() {

	angkaSatu := []int{1, 2, 3, 4, 5}

	for i := len(angkaSatu) - 1; i >= 0; i-- {
		fmt.Print(" ", angkaSatu[i])
	}

}
