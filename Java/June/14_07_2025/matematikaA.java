import java.util.Scanner;

public class matematikaA {
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        System.out.println("Program Matematika A");

        System.out.print("Masukkan bilangan bulat positf x: ");
        int x = input.nextInt();

        System.out.print("Masukan bilangan bulat positif y: ");
        int y = input.nextInt();

        // Validasi input
        if (x < 0 || y < 0) {
            System.out.println("Bilangan harus positif.");
            return;
        }

        // Menghitung hasil perhitungan f(x,y)
        double f = (1.0 / (3 * (x * x) + 10)) + ((10 * y) + 7);
        System.out.println("Hasil dari f(x,y) = 1/(3x^2 + 10) + (10y + 7) adalah: " + f);
    }
}
