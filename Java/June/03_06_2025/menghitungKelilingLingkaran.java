import java.util.Scanner;

public class menghitungKelilingLingkaran {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        System.out.println("Program Menghitung Keliling Lingkaran");

        System.out.print("Masukkan jari-jari lingkaran: ");
        double jariJari = input.nextDouble();

        double pi = 3.14;

        // Menghitung keliling lingkaran
        double keliling = 2 * pi * jariJari;
        System.out.println("Keliling lingkaran dengan jari-jari " + jariJari + " adalah " + keliling + ".");

    }
}
