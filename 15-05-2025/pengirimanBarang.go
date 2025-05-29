package main

func main() {

	beratPaket := []int{7, 15, 9, 13, 11, 5, 20}
	uangTambah, a := 0, 0

	for i := 0; i < len(beratPaket); i++ {
		if beratPaket[i] > 10 {
			a++
			uangTambah += 5000
		}
	}

	println("Jumlah paket yang terkena biaya tambahan:", a, "Uang tambahan:", uangTambah)

}
