package main

import "fmt"

func main() {

	var uang, pecahan int
	fmt.Print("Masukkan jumlah uang: ")
	fmt.Scan(&uang)

	for i := 1; i <= uang; i++ {
		if uang >= 100000 {
			pecahan = uang / 100000
			fmt.Printf("%d lembar Rp. 100.000\n", pecahan)
			uang = uang % 100000
		} else if uang >= 50000 {
			pecahan = uang / 50000
			fmt.Printf("%d lembar Rp. 50.000\n", pecahan)
			uang = uang % 50000
		} else if uang >= 20000 {
			pecahan = uang / 20000
			fmt.Printf("%d lembar Rp. 20.000\n", pecahan)
			uang = uang % 20000
		} else if uang >= 10000 {
			pecahan = uang / 10000
			fmt.Printf("%d lembar Rp. 10.000\n", pecahan)
			uang = uang % 10000
		} else if uang >= 5000 {
			pecahan = uang / 5000
			fmt.Printf("%d lembar Rp. 5.000\n", pecahan)
			uang = uang % 5000
		} else if uang >= 2000 {
			pecahan = uang / 2000
			fmt.Printf("%d lembar Rp. 2.000\n", pecahan)
			uang = uang % 2000
		} else if uang >= 1000 {
			pecahan = uang / 1000
			fmt.Printf("%d lembar Rp. 1.000\n", pecahan)
			uang = uang % 1000
		}
		if uang == 0 {
			break
		}
	}
}
