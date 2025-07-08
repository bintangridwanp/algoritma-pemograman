import java.util.Scanner;

public class diskonBelanja {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Diskon Belanja");

        System.out.print("Masukkan total belanja Anda: ");
        double totalBelanja = input.nextDouble();

        // Validasi input
        if (totalBelanja < 0) {
            System.out.println("Total belanja tidak boleh negatif.");
            return;
        }

        // Menghitung diskon
        double diskon;
        if (totalBelanja >= 100000) {
            diskon = totalBelanja * 10 / 100; // Diskon 10%
        } else {
            diskon = 0;
        }

        // Menghitung total setelah diskon
        double totalSetelahDiskon = totalBelanja - diskon;
        System.out.println("Total belanja Anda: Rp " + totalBelanja);
        System.out.println("Diskon yang diberikan: Rp " + diskon);
        System.out.println("Total yang harus dibayar setelah diskon: Rp " + totalSetelahDiskon);
    }
}
