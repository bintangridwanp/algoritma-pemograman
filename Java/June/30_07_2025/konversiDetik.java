import java.util.Scanner;

public class konversiDetik {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Konversi Detik ke Jam, Menit, dan Detik");

        System.out.print("Masukkan jumlah detik: ");
        int totalDetik = input.nextInt();

        // Hitung jumlah jam, menit, dan detik
        int jam = totalDetik / 3600;
        int sisaDetik = totalDetik % 3600;
        int menit = sisaDetik / 60;
        sisaDetik = sisaDetik % 60;
        int detik = sisaDetik % 60;

        // Tampilkan hasil konversi
        System.out.println("Konversi dari " + totalDetik + " detik adalah:");
        System.out.println(jam + " jam, " + menit + " menit, dan " + detik + " detik.");
    }
}
