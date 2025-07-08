import java.util.Scanner;

public class menampilkanDeretAngka {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menampilkan Deret Angka");

        System.out.print ("Masukkan deret angka 1 - n yang ingin ditampilkan: ");
        int jumlahAngka = input.nextInt();

        // Menampilkan deret angka dari 1 sampai n
        for (int i = 1; i <= jumlahAngka; i++) {
            System.out.print(i + " ");
        }
    }
}
