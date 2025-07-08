import java.util.Scanner;

public class menampilkanBilanganPrima {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menampilkan Bilangan Prima");

        System.out.println("Masukan bilangan bulat positif untuk menampilkan bilangan prima: ");
        int angka = input.nextInt();

        for (int i = 2; i <= angka; i++) {
            boolean isPrima = true;

            // Check divisibility up to i / 2
            for (int j = 2; j <= i / 2; j++) {
                if (i % j == 0) {
                    isPrima = false;
                    break;
                }
            }

            if (isPrima) {
                System.out.print(i + " ");
            }
        }
    }
}