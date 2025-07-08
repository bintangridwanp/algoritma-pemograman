import java.util.Scanner;

public class menghitungRataNilai {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Rata-Rata Nilai Mahasiswa");

        System.out.print("Masukkan nilai Matemetika: ");
        double nilaiMatematika = input.nextDouble();

        System.out.print("Masukkan nilai Bahasa Indonesia: ");
        double nilaiBahasaIndonesia = input.nextDouble();

        System.out.print("Masukkan nilai Bahasa Inggris: ");
        double nilaiBahasaInggris = input.nextDouble();

        // Validasi input
        if (nilaiMatematika < 0 || nilaiBahasaIndonesia < 0 || nilaiBahasaInggris < 0) {
            System.out.println("Nilai tidak boleh negatif.");
            return;
        }

        // Menghitung rata-rata nilai
        double rataRata = (nilaiMatematika + nilaiBahasaIndonesia + nilaiBahasaInggris) / 3;
        System.out.printf("Rata-rata nilai mahasiswa adalah: %.2f%n", rataRata);

        // Validasi lulus
        if (rataRata >= 70) {
            System.out.println("Mahasiswa dinyatakan LULUS.");
        } else {
            System.out.println("Mahasiswa dinyatakan TIDAK LULUS.");
        }


    }
}
