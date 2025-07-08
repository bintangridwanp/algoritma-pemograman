import java.util.Scanner;
public class kalkulatorPajakPenjualan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Kalkulator Pajak Penjualan");

        System.out.println("Masukkan total penjualan: ");
        double totalPenjualan = input.nextDouble();

        // Validasi input
        if (totalPenjualan < 0) {
            System.out.println("Total penjualan tidak boleh negatif.");
            return;
        }

        // Menghitung pajak penjualan
        if (totalPenjualan > 200) {
            double pajakPenjualan = totalPenjualan * 0.1;
            System.out.printf("Pajak penjualan yang harus dibayar adalah: %.2f%n", pajakPenjualan);
        } else {
            System.out.println("Tidak ada pajak penjualan yang dikenakan.");
        }
    }
}
