import java.util.Scanner;

public class cekPalindrom {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cek Palindrom");

        System.out.print("Masukkan sebuah kata: ");
        String kata = input.nextLine();

        boolean isPalindrom = false;
        for (int i = 0; i < kata.length() / 2; i++) {
            if (kata.charAt(i) ==  kata.charAt(kata.length() - 1 - i)) {
                isPalindrom = true;
                break;
            }
        }

        if (isPalindrom) {
            System.out.println("Kata '" + kata + "' adalah palindrom.");
        } else {
            System.out.println("Kata '" + kata + "' bukan palindrom.");
        }

    }
}
