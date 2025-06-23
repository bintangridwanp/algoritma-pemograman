import java.util.Scanner;

public class menghitungTotalBelanja {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Total Belanja");

        System.out.print("Masukkan harga barang: ");
        double hargaBarang = input.nextDouble();

        System.out.print("Masukkan jumlah barang yang dibeli: ");
        int jumlahBarang = input.nextInt();

        // Menghitung total belanja
        double totalBelanja = hargaBarang * jumlahBarang;
        System.out.println("Total belanja untuk " + jumlahBarang + " barang dengan harga " + hargaBarang + " adalah " + totalBelanja + ".");

        input.close();
    }
}
