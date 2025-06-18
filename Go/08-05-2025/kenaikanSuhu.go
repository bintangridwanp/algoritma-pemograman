package main

func main() {

	suhu := []int{30, 32, 31, 33, 35, 34, 36}
	kenaikanSuhu := 0

	for i := 0; i < len(suhu)-1; i++ {
		if suhu[i+1] > suhu[i] {
			kenaikanSuhu++
		}
	}

	println("Jumlah kenaikan suhu:", kenaikanSuhu)

}
