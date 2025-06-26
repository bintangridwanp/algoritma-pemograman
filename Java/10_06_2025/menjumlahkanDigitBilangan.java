import java.util.Scanner;

public class menjumlahkanDigitBilangan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menjumlahkan Digit Bilangan");

        System.out.print("Masukkan bilangan bulat positif: ");
        int angka = input.nextInt();

        while (angka > 0) {
            int digit = angka % 10; // Mendapatkan digit terakhir
            angka = angka / 10; // Menghapus digit terakhir dari angka
            System.out.print(digit + " "); // Menampilkan digit
        }
    }
}
