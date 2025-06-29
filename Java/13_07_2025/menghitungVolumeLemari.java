import java.util.Scanner;

public class menghitungVolumeLemari {
    public static void main(String[] args) {
       Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Volume Lemari");

        System.out.println("Masukkan panjang, lebar, dan tinggi lemari dalam satuan cm:");

        System.out.print("Panjang: ");
        double panjang = input.nextDouble();

        System.out.print("Lebar: ");
        double lebar = input.nextDouble();

        System.out.print("Masukkan tinggi: ");
        double tinggi = input.nextDouble();

        // Validasi input
        if (panjang <= 0 || lebar <= 0 || tinggi <= 0) {
            System.out.println("Panjang, lebar, dan tinggi harus lebih besar dari nol.");
            return;
        }

        // Menghitung volume lemari
        double volume = panjang * lebar * tinggi;

        System.out.println("Volume lemari adalah: " + volume + " cm³");

    }
}
