import java.util.Scanner;

public class menampilkanKarakterTerakhir {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menampilkan Karakter Terakhir dari String");
        System.out.print("Masukkan sebuah kata: ");
        String kata = input.nextLine();

        // Memeriksa apakah string tidak kosong
        if (kata.length() > 0) {
            char karakterTerakhir = kata.charAt(kata.length() - 1);
            System.out.println("Karakter terakhir dari kata '" + kata + "' adalah: " + karakterTerakhir);
        } else {
            System.out.println("Kata tidak boleh kosong.");
        }
    }
}
