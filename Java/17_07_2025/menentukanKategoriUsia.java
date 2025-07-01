import java.util.Scanner;

public class menentukanKategoriUsia {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menentukan Kategori Usia");

        System.out.print("Masukkan usia Anda: ");
        int usia = input.nextInt();

        // Menentukan kategori usia
        if (usia < 0) {
            System.out.println("Usia tidak boleh negatif.");
        } else if (usia <= 12) {
            System.out.println("Anda adalah anak-anak.");
        } else if (usia <= 17) {
            System.out.println("Anda adalah remaja.");
        } else if (usia <= 59) {
            System.out.println("Anda adalah dewasa.");
        } else {
            System.out.println("Anda adalah lansia.");
        }
    }
}
