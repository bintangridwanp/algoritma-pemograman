package main

import "fmt"

func main1() {

	kata := ("Golang");

	hasil  := " "
	

	for i:= len(kata) - 1; i >= 0; i-- {
		hasil = hasil  + string(kata[i])
	}

	fmt.Println("Hasil balik kata adalah ",hasil);
}