import java.util.Scanner;

public class kalkulatorTabunganDenganBunga {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Kalkulator Tabungan dengan Bunga");

        System.out.print("Masukkan jumlah tabungan awal: ");
        double jumlahTabungan = input.nextDouble();

        System.out.print("Masukkan suku bunga tahunan (dalam persen): ");
        double sukuBunga = input.nextDouble();

        System.out.print("Masukkan jumlah tahun: ");
        int jumlahTahun = input.nextInt();

        // Validasi input
        if (jumlahTabungan < 0 || sukuBunga < 0 || jumlahTahun < 0) {
            System.out.println("Jumlah tabungan, suku bunga, dan jumlah tahun harus lebih besar atau sama dengan nol.");
            return;
        }

        double totalAkhir = jumlahTabungan;
        for (int i = 0; i < jumlahTahun; i++) {
            totalAkhir = totalAkhir * (1 + sukuBunga / 100);
        }
        double totalBunga = totalAkhir - jumlahTabungan;

        System.out.println("Total tabungan setelah " + jumlahTahun + " tahun adalah: Rp " + totalAkhir);
        System.out.println("Total bunga yang diperoleh adalah: Rp " + totalBunga);
    }
}