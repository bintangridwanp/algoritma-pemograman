import java.util.Scanner;

public class cacahUang {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cacah Uang");

        System.out.print("Masukkan jumlah uang yang akan dicacah (dalam rupiah): ");
        int jumlahUang = input.nextInt();

        // Validasi input
        if (jumlahUang < 0) {
            System.out.println("Jumlah uang tidak boleh negatif.");
        } else {

            int sepuluhRibu = jumlahUang / 10000;
            jumlahUang %= 10000;
            int limaRibu = jumlahUang / 5000;
            jumlahUang %= 5000;
            int seribu = jumlahUang / 1000;
            jumlahUang %= 1000;

            System.out.println("Jumlah uang yang dicacah:");
            System.out.println("Rp 10.000: " + sepuluhRibu + " lembar");
            System.out.println("Rp 5.000: " + limaRibu + " lembar");
            System.out.println("Rp 1.000: " + seribu + " lembar");

            System.out.println("Jumlah uang yang tersisa: Rp " + jumlahUang);
        }
    }
}
