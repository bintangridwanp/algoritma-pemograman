package main

func main() {

	nilaiUjian := []int{78, 85, 90, 66, 92, 70, 88, 91, 60, 73}
	posisiTertinggi := nilaiUjian[0]
	posisiTerendah := nilaiUjian[0]
	a, b := 0, 0

	for i := 1; i < len(nilaiUjian); i++ {
		if posisiTerendah < nilaiUjian[i] {
			posisiTerendah = nilaiUjian[i]
			a = i
		}
		if posisiTertinggi > nilaiUjian[i] {
			posisiTertinggi = nilaiUjian[i]
			b = i
		}
	}

	println("Posisi terendah:", b, "dengan nilai", posisiTertinggi)
	println("Posisi tertinggi:", a, "dengan nilai", posisiTerendah)

}
