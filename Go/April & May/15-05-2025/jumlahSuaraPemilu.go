package main

func main() {

	dataSuara := []rune{'A', 'B', 'A', 'C', 'B', 'A', 'C', 'C', 'B', 'B'}
	a, b, c := 0, 0, 0

	for i := 0; i < len(dataSuara); i++ {
		if dataSuara[i] == 'A' {
			a++
		} else if dataSuara[i] == 'B' {
			b++
		} else if dataSuara[i] == 'C' {
			c++
		}
	}

	if a > b && a > c {
		println("Calon A menang dengan jumlah suara:", a)
	} else if b > a && b > c {
		println("Calon B menang dengan jumlah suara:", b)
	} else if c > a && c > b {
		println("Calon C menang dengan jumlah suara:", c)
	}

}
