import java.util.Scanner;

public class menjumlahkanDigitBilangan {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menjumlahkan Digit Bilangan");

        System.out.print("Masukkan bilangan bulat positif: ");
        int angka = input.nextInt();

        int hasil = 0;
        while (angka > 0) {
            int digit = angka % 10;
            angka = angka / 10;
            hasil = hasil + digit;
        }

        System.out.println("Jumlah dari digit bilangan tersebut adalah " + hasil + ".");
    }
}
