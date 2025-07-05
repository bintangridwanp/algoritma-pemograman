import java.util.Scanner;

public class menentukanKarakterTerbanyak {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menentukan Karakter Terbanyak");
        System.out.print("Masukkan sebuah kata: ");
        String kata = input.nextLine();

        char karakterTerbanyak = kata.charAt(0);
        int maxFrekuensi = 0;

        for (int i = 0; i < kata.length(); i++) {
            char karakter = kata.charAt(i);
            int frekuensi = 0;

            // Hitung frekuensi karakter ini
            for (int j = 0; j < kata.length(); j++) {
                if (kata.charAt(j) == karakter) {
                    frekuensi++;
                }
            }

            // Update jika frekuensi lebih besar
            if (frekuensi > maxFrekuensi) {
                maxFrekuensi = frekuensi;
                karakterTerbanyak = karakter;
            }
        }

        System.out.println("Karakter terbanyak dalam kata '" + kata + "' adalah '" + karakterTerbanyak + "' dengan frekuensi " + maxFrekuensi + ".");
    }
}