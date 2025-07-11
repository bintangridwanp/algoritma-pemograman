import java.util.Scanner;

public class persegiPanjang {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Menghitung Luas dan Keliling Persegi Panjang");

        System.out.print("Masukkan panjang: ");
        double panjang = input.nextDouble();

        System.out.print("Masukkan lebar: ");
        double lebar = input.nextDouble();

        // validasi input
        if (panjang <= 0 || lebar <= 0) {
            System.out.println("Panjang dan lebar harus lebih besar dari 0.");
        } else {
            // Hitung luas dan keliling
            double luas = panjang * lebar;
            double keliling = 2 * (panjang + lebar);

            // Tampilkan hasil
            System.out.printf("Luas Persegi Panjang: %.2f\n", luas);
            System.out.printf("Keliling Persegi Panjang: %.2f\n", keliling);
        }

    }
}
