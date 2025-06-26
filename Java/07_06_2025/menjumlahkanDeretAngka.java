import java.util.Scanner;

public class menjumlahkanDeretAngka {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menjumlahkan Deret Angka");

        System.out.print("Masukkan jumlah angka yang ingin dijumlahkan: ");
        int jumlahAngka = input.nextInt();

        int total = 0;
        for (int i = 0; i <= jumlahAngka; i++) {
            total = total + i;
        }

        System.out.println("Jumlah dari deret angka 0 sampai n adalah " + total + ".");
    }
}
