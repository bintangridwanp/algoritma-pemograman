package main

func main() {

	naik := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//turun := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	//acak := []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 10}

	a := false
	//b := false
	//c := false

	for i := 0; i < len(naik)-1; i++ {
		a = false
		//b := false
		//c := false
		if naik[i] < naik[i+1] {
			a = true
		} else if naik[i] > naik[i+1] {
			println("Turun")
		} else {
			println("Tetap")
		}
	}

	if a == true {
		println("Pola Naik")
	} else {
		println("Kurang jelas")
	}

}
