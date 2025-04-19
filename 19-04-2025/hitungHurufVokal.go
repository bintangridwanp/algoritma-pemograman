package main

import "fmt"

func main2() {

    kalimat := "Saya suka belajar Golang"

    hasil := 0

    for i := 0; i < len(kalimat); i++ {
        if kalimat[i] == 'a' || kalimat[i] == 'i' || kalimat[i] == 'u' || kalimat[i] == 'e' || kalimat[i] == 'o' {
            hasil = hasil + 1
        } else if kalimat[i] == 'A' || kalimat[i] == 'I' || kalimat[i] == 'U' || kalimat[i] == 'E' || kalimat[i] == 'O' {
            hasil = hasil + 1
        }
    }

    fmt.Println("Jumlah huruf vokal adalah", hasil)
}