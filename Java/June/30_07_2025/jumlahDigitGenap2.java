import java.util.Scanner;

public class jumlahDigitGenap2 {
    public static void main(String[] args) {
     Scanner input = new Scanner(System.in);

        System.out.println("Program Jumlah Digit Genap Dalam Bilangan");

        System.out.print("Masukkan sebuah bilangan: ");
        int bilangan = input.nextInt();

        int jumlahDigitGenap = 0;

        while (bilangan != 0 ) {
            int digit = bilangan % 10;

            if (digit % 2 == 0) {
                jumlahDigitGenap = jumlahDigitGenap + digit;
            }

            bilangan /= 10;
        }

        System.out.println("Jumlah digit genap dalam bilangan tersebut adalah: " + jumlahDigitGenap);

    }
}
