import java.util.Scanner;

public class menghitungAngkaKalimat {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Jumlah kata dalam Kalimat");

        System.out.println("Masukkan sebuah kalimat:");
        String kalimat = input.nextLine();

        // Validasi input
        if (kalimat.isEmpty()) {
            System.out.println("Kalimat tidak boleh kosong.");
            return;
        }

        // Menghitung jumlah kata
        int jumlahKata = 0;

        for (int i = 0; i < kalimat.length(); i++) {
            char c = kalimat.charAt(i);
            // Cek apakah karakter adalah spasi
            if (c == ' ') {
                jumlahKata++;
            }
        }

        System.out.println("Jumlah kata dalam kalimat adalah: " + (jumlahKata + 1));


    }
}
