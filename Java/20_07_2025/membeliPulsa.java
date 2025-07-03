import java.util.Scanner;

public class membeliPulsa {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Membeli Pulsa");
        System.out.println("Masukkan nominal pulsa yang ingin dibeli:");
        int nominalPulsa = input.nextInt();

        if (nominalPulsa < 5000) {
            System.out.println("Nominal pulsa terlalu kecil. Minimal pembelian adalah Rp 5.000.");
            return;
        }

        int sisa = nominalPulsa;

        int jumlah100k = sisa / 100000;
        sisa = sisa % 100000;

        int jumlah50k = sisa / 50000;
        sisa = sisa % 50000;

        int jumlah20k = sisa / 20000;
        sisa = sisa % 20000;

        int jumlah10k = sisa / 10000;
        sisa = sisa % 10000;

        int jumlah5k = sisa / 5000;
        sisa = sisa % 5000;

        System.out.println("Nominal pulsa yang dibeli: Rp " + nominalPulsa);
        System.out.println("Rincian pecahan:");
        if (jumlah100k > 0) System.out.println("Rp 100000 x " + jumlah100k);
        if (jumlah50k > 0) System.out.println("Rp 50000 x " + jumlah50k);
        if (jumlah20k > 0) System.out.println("Rp 20000 x " + jumlah20k);
        if (jumlah10k > 0) System.out.println("Rp 10000 x " + jumlah10k);
        if (jumlah5k > 0) System.out.println("Rp 5000 x " + jumlah5k);
    }
}