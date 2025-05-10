package main

import "fmt"

func main() {

	waktuKirim := [...]int{25, 35, 40, 28, 32, 20, 33, 45}
	var batasWaktu, sum int
	batasWaktu = 30

	for i := 0; i < len(waktuKirim); i++ {
		if waktuKirim[i] > batasWaktu {
			sum++
		}
	}
	fmt.Println(sum, "Pengantar melebihi waktu", batasWaktu, "menit")
}
