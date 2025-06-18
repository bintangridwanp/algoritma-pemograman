package main

func main() {

	targetHarian := []int{98, 110, 105, 99, 120, 101}
	hari := 0

	for h := 0; h < len(targetHarian); h++ {
		if targetHarian[h] >= 100 {
			hari++
		}
	}

	println("Jumlah hari produksi di atas target:", hari)
}
