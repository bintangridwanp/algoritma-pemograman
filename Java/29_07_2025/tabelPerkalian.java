import java.util.Scanner;

public class tabelPerkalian {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Tabel Perkalian");

        System.out.print("Masukkan angka untuk membuat tabel perkalian: ");
        int angka = input.nextInt();

        // Validasi input angka
        if (angka < 1 || angka > 10) {
            System.out.println("Angka tidak valid. Silakan masukkan angka antara 1 dan 10.");
        } else {
            for (int i = 1; i <= angka; i++) {
                int hasil = i * i;
                System.out.println("Hasil perkalian " + i + " x " + i + " = " + hasil);

            }
        }
    }
}
