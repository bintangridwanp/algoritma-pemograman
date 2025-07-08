import java.util.Scanner;

public class cekAnagram {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Cek Anagram Dua Kata");

        System.out.print("Masukkan kata pertama: ");
        String kata1 = input.nextLine().replaceAll("\\s+", "").toLowerCase();

        System.out.print("Masukkan kata kedua: ");
        String kata2 = input.nextLine().replaceAll("\\s+", "").toLowerCase();

        if (kata1.isEmpty() || kata2.isEmpty()) {
            System.out.println("Kata tidak boleh kosong.");
            return;
        }

        if (kata1.length() != kata2.length()) {
            System.out.println("Kata-kata tersebut bukan anagram.");
            return;
        }

        int[] frekuensi = new int[26];
        for (int i = 0; i < kata1.length(); i++) {
            frekuensi[kata1.charAt(i) - 'a']++;
            frekuensi[kata2.charAt(i) - 'a']--;
        }

        boolean anagram = true;
        for (int f : frekuensi) {
            if (f != 0) {
                anagram = false;
                break;
            }
        }

        if (anagram) {
            System.out.println("Kata-kata tersebut anagram.");
        } else {
            System.out.println("Kata-kata tersebut bukan anagram.");
        }
    }
}