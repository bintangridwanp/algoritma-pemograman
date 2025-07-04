import java.util.Scanner;

public class cekBilanganAmstrong {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cek Bilangan Armstrong");
        System.out.print("Masukkan bilangan bulat positif untuk mengecek apakah itu bilangan Armstrong: ");
        int bilangan = input.nextInt();

        if (bilangan < 0) {
            System.out.println("Bilangan harus positif.");
            return;
        }

        int temp = bilangan;
        int jumlahDigit = 0;
        while (temp != 0) {
            jumlahDigit++;
            temp /= 10;
        }

        temp = bilangan;
        int sum = 0;
        while (temp != 0) {
            int digit = temp % 10;
            int pangkat = 1;
            for (int i = 0; i < jumlahDigit; i++) {
                pangkat *= digit;
            }
            sum += pangkat;
            temp /= 10;
        }

        if (sum == bilangan) {
            System.out.println(bilangan + " adalah bilangan Armstrong.");
        } else {
            System.out.println(bilangan + " bukan bilangan Armstrong.");
        }
    }
}