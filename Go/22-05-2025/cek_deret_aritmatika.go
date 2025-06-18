package main

func main() {

	angka := []int{3, 6, 9, 12}
	hasil := true

	for i := 0; i < len(angka)-1; i++ {
		if angka[i+1]-angka[i] != angka[1]-angka[0] {
			hasil = false
		}
	}
	
	if hasil {
		println("Deret Aritmatika")
	} else {
		println("Bukan Deret Aritmatika")
	}

}
