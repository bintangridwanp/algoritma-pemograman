package main

import "fmt"

func main() {
	var input int

	for {
		println("Masukkan total belanja produk (atau masukkan 0 untuk keluar):")
		fmt.Scanln(&input)

		if input <= 0 {
			println("Program selesai.")
			break
		}

		var totalBelanja int
		for i := 1; i <= input; i++ {
			var harga int
			println("Masukkan harga produk ke-", i, ":")
			fmt.Scanln(&harga)
			totalBelanja += harga
		}

		println("Total belanja produk adalah:", totalBelanja)
	}

}
