import java.util.Scanner;

public class matematikaB {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Matematika B");

        System.out.print("Masukkan bilangan positf x untuk melakukan perhitungan: ");
        double x = input.nextDouble();

        // Validasi input
        if (x < 0) {
            System.out.println("Bilangan harus positif.");
            return;
        }

        // Menghitung hasil perhitungan f(x)
        double f = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * (x * x)) + 4);
        System.out.println("Nilai dari f(x) = (x^3 + 3x) / (x^4 - 3x^2 + 4) adalah: " + f);

    }
}
