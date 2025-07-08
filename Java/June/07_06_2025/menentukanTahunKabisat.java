import java.util.Scanner;

public class menentukanTahunKabisat {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menentukan Tahun Kabisat");

        System.out.println("Masukkan tahun yang ingin diperiksa: ");
        int tahun = input.nextInt();

        // Menentukan apakah tahun kabisat
        boolean isKabisat = false;
        if (tahun % 4 == 0) {
            if (tahun % 100 != 0 || tahun % 400 == 0) {
                isKabisat = true;
            }
        }

        if (isKabisat) {
            System.out.println(tahun + " adalah tahun kabisat.");
        } else {
            System.out.println(tahun + " bukan tahun kabisat.");
        }

    }
}
