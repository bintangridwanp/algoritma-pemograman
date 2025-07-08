import java.util.Scanner;
public class menghitungBiayaListrik {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Biaya Listrik");

        System.out.println("Masukkan jumlah pemakaian listrik (dalam kWh): ");
        double pemakaianListrik = input.nextDouble();

        double totalBiaya = 0;

        for (int i = 1; i <= pemakaianListrik; i++) {
            if (i <= 100) {
                totalBiaya += 1.500;
            } else  {
                totalBiaya += 2.000;
            }
        }

        System.out.println("Total biaya listrik untuk " + pemakaianListrik + " kWh adalah: Rp " + totalBiaya);
        input.close();
    }
}
