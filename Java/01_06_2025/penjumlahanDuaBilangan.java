import java.util.Scanner;

public class penjumlahanDuaBilangan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.print("Masukkan bilangan pertama: ");
        int bilanganPertama = input.nextInt();

        System.out.print("Masukkan bilangan kedua: ");
        int bilanganKedua = input.nextInt();

        int hasilPenjumlahan = bilanganPertama + bilanganKedua;
        System.out.println("Hasil penjumlahan: " + hasilPenjumlahan);
        input.close();
    }
}
