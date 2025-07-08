package main

import "fmt"

func main() {

	nilaiLari := []int{11, 13, 12, 14, 10, 12, 15, 11, 13, 9}
	atletLariLolos := 0

	for i := 0; i < len(nilaiLari); i++ {
		if nilaiLari[i] <= 12 {
			atletLariLolos++
		}
	}

	fmt.Println("Jumlah atlet yang lolos:", atletLariLolos)

}
