import java.util.Scanner;

public class transaksiBelanja {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Transaksi Belanja");

        System.out.print("Masukkan jumlah barang yang dibeli: ");
        int jumlahBarang = input.nextInt();

        double totalPembayaran = 0;

        for (int i = 0; i < jumlahBarang; i++) {
            System.out.println("Barang ke-" + (i + 1));

            System.out.print("Masukkan harga barang: ");
            double hargaBarang = input.nextDouble();

            System.out.print("Masukkan jumlah barang: ");
            int jumlahBarangDibeli = input.nextInt();

            // Validasi input
            if (hargaBarang < 0 || jumlahBarangDibeli < 0) {
                System.out.println("Harga barang dan jumlah barang harus lebih dari atau sama dengan 0.");
                i--;
                continue;
            }
            double totalHarga = hargaBarang * jumlahBarangDibeli;
            System.out.printf("Total harga untuk barang ke-%d adalah: Rp %.2f%n", (i + 1), totalHarga);

             totalPembayaran += totalHarga;
        }
        System.out.println("Total pembayaran untuk semua barang adalah: Rp " + totalPembayaran);
    }
}
