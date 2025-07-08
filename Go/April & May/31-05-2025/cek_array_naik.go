package main

func main() {

	angka := []int{1, 2, 2, 3, 4}
	hasil := false

	for i := 0; i < len(angka); i++ {
		if angka[i] >= angka[i+1] {
			hasil = true
			break
		}
	}

	if hasil {
		println("Array Naik")
	} else {
		println("Bukan Array Naik")
	}

}
