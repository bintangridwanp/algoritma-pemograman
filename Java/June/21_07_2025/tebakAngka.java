import java.util.Scanner;

public class tebakAngka {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Program Tebak Angka");
        System.out.println("Tebak angka antara 1 sampai 10");

        int angkaTebakan = 7;
        int tebakan = 0;


        while (tebakan != angkaTebakan) {
            System.out.print("Masukkan tebakan Anda: ");
            tebakan = input.nextInt();

            if (tebakan < angkaTebakan) {
                System.out.println("Tebakan Anda terlalu rendah. Coba lagi!");
            } else if (tebakan > angkaTebakan) {
                System.out.println("Tebakan Anda terlalu tinggi. Coba lagi!");
            } else {
                System.out.println("Selamat! Tebakan Anda benar.");
            }
        }
    }
}
