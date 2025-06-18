package main

func main() {

	waktuPengiriman := []int{45, 30, 50, 25, 35, 40, 60, 28}
	tersingkat := waktuPengiriman[0]
	a := 0

	for i := 1; i < len(waktuPengiriman); i++ {
		if waktuPengiriman[i] < tersingkat {
			tersingkat = waktuPengiriman[i]
			a = i + 1
		}
	}

	println("Waktu pengiriman tersingkat adalah:", tersingkat, "menit pada pengiriman ke-", a)

}
