import java.util.Scanner;

public class jumlahDigitGenap {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Jumlah Digit Genap");

        System.out.print("Masukkan bilangan bulat positif: ");
        int bilangan = input.nextInt();

        int jumlahDigitGenap = 0;
        while (bilangan > 0) {
            int digit = bilangan % 10;
            if (digit % 2 == 0) {
                jumlahDigitGenap++;
            }
            bilangan /= 10;
        }
        System.out.println("Jumlah digit genap: " + jumlahDigitGenap);
    }
}
