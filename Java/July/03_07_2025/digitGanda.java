import java.util.Scanner;

public class digitGanda {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Digit Ganda");

        System.out.print("Masukkan angka: ");
        int angka = input.nextInt();

        if (angka < 0) {
            System.out.println("Angka harus lebih besar atau sama dengan 0.");
        } else {
            String angkaStr = Integer.toString(angka);
            for (int i = 0; i < angkaStr.length(); i++) {
                char digit = angkaStr.charAt(i);
                System.out.print(digit);
                System.out.print(digit);
            }
        }
    }
}