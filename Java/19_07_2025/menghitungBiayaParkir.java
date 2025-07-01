import java.util.Scanner;

public class menghitungBiayaParkir {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Biaya Parkir");

        System.out.print("Masukkan lama parkir dalam jam: ");
        int lamaParkir = input.nextInt();

        // Validasi input
        if (lamaParkir < 1 || lamaParkir > 100) {
            System.out.println("Lama parkir harus antara 1 dan 100 jam.");
            return;
        }

        int biaya = 0;
        for (int i = 1; i <= lamaParkir; i++) {
            if (i <= 2) {
                biaya = biaya + 3000;
            } else {
                biaya = biaya + 1000;
            }
            System.out.println("Biaya parkir untuk " + i + " jam: Rp" + biaya);
        }

        System.out.println("Total biaya parkir untuk " + lamaParkir + " jam: Rp" + biaya);

        System.out.println("Terima kasih telah menggunakan layanan parkir kami!");
    }
}
