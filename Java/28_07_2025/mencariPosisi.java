import java.util.Scanner;

public class mencariPosisi {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Proogram Mencari Posisi");

        System.out.print("Masukkan jarak awal (dalam meter): ");
        double jarakAwal = input.nextDouble();

        System.out.print("Masukkan kecepatan (dalam m/s): ");
        double kecepatan = input.nextDouble();

        System.out.print("Masukkan waktu (dalam detik): ");
        double waktu = input.nextDouble();

        // Menghitung posisi
        double posisi = jarakAwal + (kecepatan * waktu);

        System.out.println("Posisi objek setelah " + waktu + " detik adalah: " + posisi + " meter");
    }
}
