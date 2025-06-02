package main

func main() {

	daftarBelanja := []string{"gula", "beras", "minyak", "gula", "sabun", "minyak"}
	daftarUnik := []string{}

	for i := 0; i < len(daftarBelanja); i++ {
		sudahAda := false
		for j := 0; j < len(daftarUnik); j++ {
			if daftarBelanja[i] == daftarUnik[j] {
				sudahAda = true
				break
			}
		}
		if !sudahAda {
			daftarUnik = append(daftarUnik, daftarBelanja[i])
		}
	}
	for i := 0; i < len(daftarUnik); i++ {
		println(daftarUnik[i])
	}

}
