import java.util.Scanner;

public class membalikanKata {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Membalikkan Kata");

        System.out.print("Masukkan sebuah kata: ");
        String kata = input.nextLine();

        // Membalikkan kata
        String kataTerbalik = "";
        for (int i = 0; i < kata.length(); i++) {
            kataTerbalik = kata.charAt(i) + kataTerbalik; //Menambahkan karakter di depan

        }

        System.out.println("Kata '" + kata + "' setelah dibalik adalah: " + kataTerbalik);

    }
}
