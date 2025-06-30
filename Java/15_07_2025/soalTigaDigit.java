import java.util.Scanner;

public class soalTigaDigit {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.print("Masukkan bilangan bulat positif (0-999): ");
        int angka = input.nextInt();

        if (angka < 0 || angka > 999) {
            System.out.println("Bilangan harus antara 0 dan 999.");
            return;
        }

        int ratusan = angka / 100;
        int puluhan = (angka % 100) / 10;
        int satuan = angka % 10;

        System.out.println(ratusan + " " + puluhan + " " + satuan);
    }
}