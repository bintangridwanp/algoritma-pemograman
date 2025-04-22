package main

import "fmt"

func main() {
    kalimat := "Aku suka belajar golang"
    terbanyak := 0
    karakterTerbanyak := ' '

    // Loop untuk menghitung jumlah kemunculan setiap huruf
    for i := 0; i < len(kalimat); i++ {
        // Lewati spasi
        if kalimat[i] == ' ' {
            continue
        }

        // Hitung jumlah kemunculan huruf
        jumlah := 0
        for j := 0; j < len(kalimat); j++ {
            if kalimat[i] == kalimat[j] {
                jumlah++
            }
        }

        // Periksa apakah huruf ini muncul lebih banyak dari sebelumnya
        if jumlah > terbanyak {
            terbanyak = jumlah
            karakterTerbanyak = rune(kalimat[i])
        }
    }

    // Tampilkan hasil
    fmt.Printf("Karakter yang paling sering muncul adalah '%c' dengan jumlah %d kali.\n", karakterTerbanyak, terbanyak)
}