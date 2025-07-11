import java.util.Scanner;

public class toko {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Toko Sederhana");

        System.out.print("Masukan harga barang 1: ");
        double hargaBarang1 = input.nextDouble();

        System.out.print("Masukan harga barang 2: ");
        double hargaBarang2 = input.nextDouble();

        System.out.print("Masukan harga barang 3: ");
        double hargaBarang3 = input.nextDouble();

        // validasi input
        if (hargaBarang1 < 0 || hargaBarang2 < 0 || hargaBarang3 < 0) {
            System.out.println("Harga barang tidak boleh negatif.");
        } else {
            // hitung keuntungngan
            double barang1 = hargaBarang1 + (hargaBarang1 * 0.05);
            double barang2 = hargaBarang2 + (hargaBarang2 * 0.05);
            double barang3 = hargaBarang3 + (hargaBarang3 * 0.05);

            System.out.println("Daftar Harga Barang Setelah Keuntungan:");
            System.out.printf("Barang 1: Rp. %.2f\n", barang1);
            System.out.printf("Barang 2: Rp. %.2f\n", barang2);
            System.out.printf("Barang 3: Rp. %.2f\n", barang3);


        }
}
}
