import java.util.Scanner;

public class menghitungUmur {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Umur");

        System.out.print("Masukan tahun lahir Anda: ");
        int tahunLahir = input.nextInt();

        int hasil = 2025 - tahunLahir;
        System.out.print("Umur Anda pada tahun 2025 adalah: " + hasil + " tahun.");


    }
}
